// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package cmd

import (
	"errors"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.nlx.io/nlx/common/cmd"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/management-api/pkg/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/basicauth"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/oidc"
	"go.nlx.io/nlx/management-api/pkg/txlogdb"
)

type oidcOptions = oidc.Options

var serveOpts struct {
	ListenAddress                string
	ConfigListenAddress          string
	PostgresDSN                  string
	DirectoryInspectionAddress   string
	DirectoryRegistrationAddress string
	TransactionLogDSN            string
	EnableBasicAuth              bool

	logoptions.LogOptions
	cmd.TLSOrgOptions
	cmd.TLSOptions
	oidcOptions
}

//nolint:gochecknoinits,funlen,gocyclo // this is the recommended way to use cobra, also a lot of flags..
func init() {
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddress, "listen-address", "", "127.0.0.1:8080", "Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.ConfigListenAddress, "config-listen-address", "", "127.0.0.1:8443", "Address for the configapi to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	serveCommand.Flags().StringVarP(&serveOpts.DirectoryInspectionAddress, "directory-inspection-address", "", "", "Address of the directory inspection API")
	serveCommand.Flags().StringVarP(&serveOpts.DirectoryRegistrationAddress, "directory-registration-address", "", "", "Address of the directory registration API")
	serveCommand.Flags().BoolVarP(&serveOpts.EnableBasicAuth, "enable-basic-auth", "", false, "Enable HTTP basic authentication and disable OIDC")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogType, "log-type", "", "live", "Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogLevel, "log-level", "", "", "Override the default loglevel as set by --log-type.")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.RootCertFile, "tls-root-cert", "", "", "Absolute or relative path to the CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.CertFile, "tls-cert", "", "", "Absolute or relative path to the cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.KeyFile, "tls-key", "", "", "Absolute or relative path to the key .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.NLXRootCert, "tls-nlx-root-cert", "", "", "Absolute or relative path to the NLX CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgCertFile, "tls-org-cert", "", "", "Absolute or relative path to the Organization cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgKeyFile, "tls-org-key", "", "", "Absolute or relative path to the Organization key .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TransactionLogDSN, "transaction-log-dsn", "", "", "Postgres DSN to the transaction log database")
	serveCommand.Flags().StringVarP(&serveOpts.oidcOptions.SecretKey, "secret-key", "", "", "Secret key that is used for signing sessions")
	serveCommand.Flags().StringVarP(&serveOpts.oidcOptions.ClientID, "oidc-client-id", "", "", "The OIDC client ID")
	serveCommand.Flags().StringVarP(&serveOpts.oidcOptions.ClientSecret, "oidc-client-secret", "", "", "The OIDC client secret")
	serveCommand.Flags().StringVarP(&serveOpts.oidcOptions.DiscoveryURL, "oidc-discovery-url", "", "", "The OIDC discovery URL")
	serveCommand.Flags().StringVarP(&serveOpts.oidcOptions.RedirectURL, "oidc-redirect-url", "", "", "The OIDC redirect URL")
	serveCommand.Flags().BoolVarP(&serveOpts.oidcOptions.SessionCookieSecure, "session-cookie-secure", "", false, "Use 'secure' cookies")

	if err := serveCommand.MarkFlagRequired("postgres-dsn"); err != nil {
		log.Fatal(err)
	}

	if err := serveCommand.MarkFlagRequired("directory-inspection-address"); err != nil {
		log.Fatal(err)
	}

	if err := serveCommand.MarkFlagRequired("directory-registration-address"); err != nil {
		log.Fatal(err)
	}

	if err := serveCommand.MarkFlagRequired("tls-nlx-root-cert"); err != nil {
		log.Fatal(err)
	}

	if err := serveCommand.MarkFlagRequired("tls-org-cert"); err != nil {
		log.Fatal(err)
	}

	if err := serveCommand.MarkFlagRequired("tls-org-key"); err != nil {
		log.Fatal(err)
	}
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start the gRPC server and HTTP API Gateway",
	Run: func(cmd *cobra.Command, args []string) {
		zapConfig := serveOpts.LogOptions.ZapConfig()
		logger, err := zapConfig.Build()
		if err != nil {
			log.Fatalf("failed to create new zap logger: %v", err)
		}

		logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
		logger = logger.With(zap.String("version", version.BuildVersion))

		if serveOpts.EnableBasicAuth {
			logger.Info("basic auth enabled, note that OIDC doesn't work")
		} else {
			if serveOpts.oidcOptions.SecretKey == "" {
				log.Fatal(errors.New("secret-key is required"))
			}

			if serveOpts.oidcOptions.ClientID == "" {
				log.Fatal(errors.New("oidc-client-id is required"))
			}

			if serveOpts.oidcOptions.ClientSecret == "" {
				log.Fatal(errors.New("oidc-client-secret is required"))
			}

			if serveOpts.oidcOptions.DiscoveryURL == "" {
				log.Fatal(errors.New("oidc-discovery-url is required"))
			}

			if serveOpts.oidcOptions.RedirectURL == "" {
				log.Fatal(errors.New("oidc-redirect-url is required"))
			}
		}

		logger.Info("starting management api", zap.String("listen-address", serveOpts.ListenAddress))

		mainProcess := process.NewProcess(logger)

		var txlogDB txlogdb.TxlogDatabase

		if len(serveOpts.TransactionLogDSN) > 0 {
			db, txlogErr := gorm.Open(postgres.Open(serveOpts.TransactionLogDSN), &gorm.Config{})
			if txlogErr != nil {
				log.Fatalf("failed to connect to the log database: %v", err)
			}

			txlogDB = &txlogdb.TxlogPostgresDatabase{DB: db}
		}

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.OrgKeyFile); errValidate != nil {
			logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", serveOpts.OrgKeyFile))
		}

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.KeyFile); errValidate != nil {
			logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", serveOpts.KeyFile))
		}

		cert, err := common_tls.NewBundleFromFiles(serveOpts.CertFile, serveOpts.KeyFile, serveOpts.RootCertFile)
		if err != nil {
			logger.Fatal("loading internal cert", zap.Error(err))
		}

		orgCert, err := common_tls.NewBundleFromFiles(serveOpts.OrgCertFile, serveOpts.OrgKeyFile, serveOpts.NLXRootCert)
		if err != nil {
			logger.Fatal("loading organization cert", zap.Error(err))
		}

		db, err := database.NewPostgresConfigDatabase(serveOpts.PostgresDSN, orgCert.Certificate().Subject.Organization[0])
		if err != nil {
			log.Fatalf("failed to connect to the database: %v", err)
		}

		auditLogger := auditlog.NewPostgresLogger(db, logger)

		var authenticator api.Authenticator

		if serveOpts.EnableBasicAuth {
			authenticator = basicauth.NewAuthenticator(db, logger)
		} else {
			authenticator = oidc.NewAuthenticator(db, auditLogger, logger, &serveOpts.oidcOptions)
		}

		a, err := api.NewAPI(
			db,
			txlogDB,
			logger,
			mainProcess,
			cert,
			orgCert,
			serveOpts.DirectoryInspectionAddress,
			serveOpts.DirectoryRegistrationAddress,
			authenticator,
			auditLogger,
		)
		if err != nil {
			logger.Fatal("cannot setup management api", zap.Error(err))
		}

		// Listen on the address provided in the serveOpts
		err = a.ListenAndServe(serveOpts.ListenAddress, serveOpts.ConfigListenAddress)
		if err != nil {
			logger.Fatal("failed to listen and serve", zap.Error(err))
		}
	},
}
