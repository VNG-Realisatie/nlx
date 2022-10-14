// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/cmd"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/strings"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/common/version"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/management-api/api"
	external_api "go.nlx.io/nlx/management-api/api/external"
)

var options struct {
	ListenAddress                   string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8443" description:"Address for the inway to listen on."`
	ListenAddressManagementAPIProxy string `long:"listen-address-management-api-proxy" env:"LISTEN_ADDRESS_MANAGEMENT_API_PROXY" default:"127.0.0.1:8444" description:"Address for the inway to listen on for management requests." required:"true"`
	Address                         string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`
	ManagementAPIAddress            string `long:"management-api-address" env:"MANAGEMENT_API_ADDRESS" description:"The address NLX Management API which will be served as a proxy." required:"true"`
	ManagementAPIProxyAddress       string `long:"management-api-proxy-address" env:"MANAGEMENT_API_PROXY_ADDRESS" description:"The address other organizations can use to reach the NLX Management API proxy." required:"true"`
	MonitoringAddress               string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"127.0.0.1:8081" description:"Address for the inway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	DirectoryRegistrationAddress    string `long:"directory-registration-address" env:"DIRECTORY_REGISTRATION_ADDRESS" description:"Address for the directory where this inway can register its services"`
	DirectoryAddress                string `long:"directory-address" env:"DIRECTORY_ADDRESS" description:"Address for the directory where this inway can register it's services"`
	DisableLogdb                    bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	PostgresDSN                     string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`
	TxLogAPIAddress                 string `long:"tx-log-api-address" env:"TX_LOG_API_ADDRESS" description:"The address of the transaction log API" required:"false"`
	Name                            string `long:"name" env:"INWAY_NAME" description:"Name of the inway. Every inway should have a unique name within the organization." required:"true"`
	AuthorizationServiceAddress     string `long:"authorization-service-address" env:"AUTHORIZATION_SERVICE_ADDRESS" description:"Address of the authorization service. If set calls will go through the authorization service before being send to the service"`
	AuthorizationCA                 string `long:"authorization-root-ca" env:"AUTHORIZATION_ROOT_CA" description:"absolute path to root CA used to verify auth service certificate"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
}

func main() {
	parseOptions()

	p := process.NewProcess()

	logger := setupLogger()

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.OrgKeyFile); errValidate != nil {
		logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", options.OrgKeyFile))
	}

	orgCert, err := common_tls.NewBundleFromFiles(options.OrgCertFile, options.OrgKeyFile, options.NLXRootCert)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	if errValidate := common_tls.VerifyPrivateKeyPermissions(options.KeyFile); errValidate != nil {
		logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", options.KeyFile))
	}

	cert, err := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
	if err != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	var txLogger transactionlog.TransactionLogger
	if options.DisableLogdb {
		txLogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		txLogger, err = setupTransactionLogger(logger, options.PostgresDSN, options.TxLogAPIAddress, cert)
		if err != nil {
			logger.Fatal("unable to setup the transaction logger", zap.Error(err))
		}
	}

	creds := credentials.NewTLS(cert.TLSConfig())

	connCtx, connCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer connCtxCancel()

	logger.Info("creating management api connection", zap.String("management api address", options.ManagementAPIAddress))

	conn, err := grpc.DialContext(connCtx, options.ManagementAPIAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		logger.Fatal("failed to setup connection to management api", zap.Error(err))
	}

	managementProxy, err := grpcproxy.New(context.TODO(), logger, options.ManagementAPIAddress, orgCert, cert)
	if err != nil {
		logger.Fatal("failed to setup management api proxy", zap.Error(err))
	}

	managementProxy.RegisterService(external_api.GetAccessRequestServiceDesc())
	managementProxy.RegisterService(external_api.GetDelegationServiceDesc())

	directoryDialCredentials := credentials.NewTLS(orgCert.TLSConfig())
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}

	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(nlxversion.NewGRPCContext(context.Background(), "inway"), 1*time.Minute)
	directoryConn, err := grpc.DialContext(directoryConnCtx, options.DirectoryAddress, directoryDialOptions...)

	defer directoryConnCtxCancel()

	if err != nil {
		logger.Fatal("failed to setup connection to directory registration api", zap.Error(err))
	}

	directoryClient := directoryapi.NewDirectoryClient(directoryConn)

	params := &inway.Params{
		Context:                         context.Background(),
		Logger:                          logger,
		Txlogger:                        txLogger,
		ManagementClient:                api.NewManagementServiceClient(conn),
		ManagementProxy:                 managementProxy,
		Name:                            options.Name,
		Address:                         options.Address,
		ManagementAPIProxyAddress:       options.ManagementAPIProxyAddress,
		MonitoringAddress:               options.MonitoringAddress,
		ListenAddressManagementAPIProxy: options.ListenAddressManagementAPIProxy,
		OrgCertBundle:                   orgCert,
		DirectoryClient:                 directoryClient,
		AuthServiceURL:                  options.AuthorizationServiceAddress,
		AuthCAPath:                      options.AuthorizationCA,
	}

	iw, err := inway.NewInway(params)
	if err != nil {
		logger.Fatal("cannot setup inway", zap.Error(err))
	}

	go func() {
		err = iw.Run(context.Background(), options.ListenAddress)
		if err != nil {
			logger.Fatal("failed to run server", zap.Error(err))
		}
	}()

	p.Wait()

	shutdown(logger, iw, txLogger)
}

func shutdown(logger *zap.Logger, iw *inway.Inway, transactionLogger transactionlog.TransactionLogger) {
	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := iw.Shutdown(gracefulCtx)
	if err != nil {
		logger.Error("failed to shutdown", zap.Error(err))
	}

	err = transactionLogger.Close()
	if err != nil {
		logger.Error("failed to close transactionLogger", zap.Error(err))
	}
}

func setupTransactionLogger(logger *zap.Logger, postgresDSN, txLogAPIAddress string, certificateBundle *common_tls.CertificateBundle) (transactionlog.TransactionLogger, error) {
	if postgresDSN != "" && txLogAPIAddress != "" {
		return nil, fmt.Errorf("cannot configure both postgresDSN and txlogAPIAddress")
	}

	if postgresDSN != "" {
		logDB, err := setupDatabase()
		if err != nil {
			return nil, err
		}

		return transactionlog.NewPostgresTransactionLogger(logger, logDB, transactionlog.DirectionIn)
	}

	if txLogAPIAddress != "" {
		return transactionlog.NewAPITransactionLogger(&transactionlog.NewAPITransactionLoggerArgs{
			Logger:       logger,
			Direction:    transactionlog.DirectionIn,
			APIAddress:   txLogAPIAddress,
			InternalCert: certificateBundle,
		})
	}

	return transactionlog.NewDiscardTransactionLogger(), nil
}

func setupLogger() *zap.Logger {
	zapConfig := options.LogOptions.ZapConfig()

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	logger.Info("starting inway", zap.String("directory-address", options.DirectoryAddress))

	return logger
}

func setupDatabase() (*sqlx.DB, error) {
	var logDB *sqlx.DB

	logDB, err := sqlx.Open("postgres", options.PostgresDSN)
	if err != nil {
		return nil, errors.Wrap(err, "could not open connection to postgres")
	}

	var (
		connMaxLifetime = 5 * time.Minute
		maxOpenConns    = 100
		maxIdleConns    = 100
	)

	logDB.SetConnMaxLifetime(connMaxLifetime)
	logDB.SetMaxOpenConns(maxOpenConns)
	logDB.SetMaxIdleConns(maxIdleConns)
	logDB.MapperFunc(strings.ToSnakeCase)

	return logDB, nil
}

func parseOptions() {
	args, err := flags.Parse(&options)
	if err != nil {
		if et, ok := err.(*flags.Error); ok {
			if et.Type == flags.ErrHelp {
				return
			}
		}

		log.Fatalf("error parsing flags: %v", err)
	}

	if len(args) > 0 {
		log.Fatalf("unexpected arguments: %v", args)
	}

	// Deprecated flags
	if options.DirectoryRegistrationAddress != "" {
		log.Println("Flag --directory-registration-address has been deprecated, please use '--directory-address' instead")
	}

	if options.DirectoryAddress == "" {
		if options.DirectoryRegistrationAddress == "" {
			log.Fatal(errors.New("directory-address is required"))
		}

		options.DirectoryAddress = options.DirectoryRegistrationAddress
	}
}
