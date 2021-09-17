// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/huandu/xstrings"
	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/cmd"
	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/inway"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/management-api/api"
	external_api "go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/txlog-db/dbversion"
)

var options struct {
	ListenAddress                string `long:"listen-address" env:"LISTEN_ADDRESS" default:"127.0.0.1:8443" description:"Address for the inway to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ListenManagementAddress      string `long:"listen-management-address" env:"LISTEN_MANAGEMENT_ADDRESS" description:"Address for the inway to listen on for management requests. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	MonitoringAddress            string `long:"monitoring-address" env:"MONITORING_ADDRESS" default:"127.0.0.1:8081" description:"Address for the inway monitoring endpoints to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	DirectoryRegistrationAddress string `long:"directory-registration-address" env:"DIRECTORY_REGISTRATION_ADDRESS" description:"Address for the directory where this inway can register it's services" required:"true"`
	DisableLogdb                 bool   `long:"disable-logdb" env:"DISABLE_LOGDB" description:"Disable logdb connections"`
	ManagementAPIAddress         string `long:"management-api-address" env:"MANAGEMENT_API_ADDRESS" description:"The address of the NLX Management API" required:"true"`
	Address                      string `long:"self-address" env:"SELF_ADDRESS" description:"The address that outways can use to reach me" required:"true"`
	PostgresDSN                  string `long:"postgres-dsn" env:"POSTGRES_DSN" description:"DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters."`
	Name                         string `long:"name" env:"INWAY_NAME" description:"Name of the inway. Every inway should have a unique name within the organization." required:"true"`

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
}

func main() {
	parseOptions()

	p := process.NewProcess()

	logger := setupLogger()
	txlogger, logDB := setupTransactionLogger(logger, options.DisableLogdb)

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

	cert, mErr := common_tls.NewBundleFromFiles(options.CertFile, options.KeyFile, options.RootCertFile)
	if mErr != nil {
		logger.Fatal("loading TLS files", zap.Error(err))
	}

	listenManagementAddress, err := defaultManagementAddress(options.ListenAddress)
	if err != nil {
		logger.Fatal("unable to create default management address", zap.Error(err))
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
		logger.Fatal("failed to setup  management api proxy", zap.Error(err))
	}

	managementProxy.RegisterService(external_api.GetAccessRequestServiceDesc())
	managementProxy.RegisterService(external_api.GetDelegationServiceDesc())

	params := &inway.Params{
		Context:                      context.Background(),
		Logger:                       logger,
		Txlogger:                     txlogger,
		ManagementClient:             api.NewManagementClient(conn),
		ManagementProxy:              managementProxy,
		Name:                         options.Name,
		Address:                      options.Address,
		MonitoringAddress:            options.MonitoringAddress,
		ListenManagementAddress:      listenManagementAddress,
		OrgCertBundle:                orgCert,
		DirectoryRegistrationAddress: options.DirectoryRegistrationAddress,
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

	shutdown(logger, iw, logDB)
}

func shutdown(logger *zap.Logger, iw *inway.Inway, logDB *sqlx.DB) {
	logger.Info("starting graceful shutdown")

	gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := iw.Shutdown(gracefulCtx)
	if err != nil {
		logger.Error("failed to shutdown", zap.Error(err))
	}

	if logDB != nil {
		err = logDB.Close()
		if err != nil {
			logger.Error("failed to close logDB", zap.Error(err))
		}
	}
}

func setupTransactionLogger(logger *zap.Logger, disabled bool) (transactionlog.TransactionLogger, *sqlx.DB) {
	if disabled {
		logger.Info("logging to transaction-log disabled")
		return transactionlog.NewDiscardTransactionLogger(), nil
	}

	logDB, err := setupDatabase(logger)
	if err != nil {
		logger.Fatal("failed to setup database", zap.Error(err))
	}

	postgresTxLogger, err := transactionlog.NewPostgresTransactionLogger(logger, logDB, transactionlog.DirectionIn)
	if err != nil {
		logger.Fatal("failed to setup transactionlog", zap.Error(err))
	}

	logger.Info("transaction logger created")

	return postgresTxLogger, logDB
}

func defaultManagementAddress(address string) (string, error) {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", err
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%v:%v", host, (intPort + 1)), nil
}

func setupLogger() *zap.Logger {
	zapConfig := options.LogOptions.ZapConfig()

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
	logger = logger.With(zap.String("version", version.BuildVersion))

	logger.Info("starting inway", zap.String("directory-registration-address", options.DirectoryRegistrationAddress))

	return logger
}

func setupDatabase(logger *zap.Logger) (*sqlx.DB, error) {
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
	logDB.MapperFunc(xstrings.ToSnakeCase)

	common_db.WaitForLatestDBVersion(logger, logDB.DB, dbversion.LatestTxlogDBVersion)

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
}
