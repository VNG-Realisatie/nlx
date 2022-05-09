// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/cmd"
	nlxhttp "go.nlx.io/nlx/common/http"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	pgadapter "go.nlx.io/nlx/directory-api/adapters/storage/postgres"
	"go.nlx.io/nlx/directory-api/http"
	"go.nlx.io/nlx/directory-api/pkg/directory"
)

var serveOpts struct {
	ListenAddress      string
	ListenAddressPlain string
	PostgresDSN        string
	TermsOfServiceURL  string

	logoptions.LogOptions

	cmd.TLSOrgOptions
	cmd.TLSOptions
}

//nolint:gochecknoinits,funlen,gocyclo // this is the recommended way to use cobra, also a lot of flags..
func init() {
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddress, "listen-address", "", "127.0.0.1:8443", "Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddressPlain, "listen-address-plain", "", "127.0.0.1:8080", "Address for the directory to listen on using plain HTTP. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.PostgresDSN, "postgres-dsn", "", "", "DSN for the postgres driver. See https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters.")
	serveCommand.Flags().StringVarP(&serveOpts.TermsOfServiceURL, "terms-of-service-url", "", "", "Link to the Terms of Service")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogType, "log-type", "", "live", "Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogLevel, "log-level", "", "", "Override the default loglevel as set by --log-type.")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.RootCertFile, "tls-root-cert", "", "", "Absolute or relative path to the CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.CertFile, "tls-cert", "", "", "Absolute or relative path to the cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.KeyFile, "tls-key", "", "", "Absolute or relative path to the key .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.NLXRootCert, "tls-nlx-root-cert", "", "", "Absolute or relative path to the NLX CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgCertFile, "tls-org-cert", "", "", "Absolute or relative path to the Organization cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgKeyFile, "tls-org-key", "", "", "Absolute or relative path to the Organization key .pem")
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
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

		p := process.NewProcess()

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.OrgKeyFile); errValidate != nil {
			logger.Warn("invalid organization key permissions", zap.Error(errValidate), zap.String("file-path", serveOpts.OrgKeyFile))
		}

		certificate, err := common_tls.NewBundleFromFiles(serveOpts.OrgCertFile, serveOpts.OrgKeyFile, serveOpts.NLXRootCert)
		if err != nil {
			logger.Fatal("loading certificate", zap.Error(err))
		}

		db, err := pgadapter.NewPostgreSQLConnection(serveOpts.PostgresDSN)
		if err != nil {
			logger.Fatal("can not create db connection:", zap.Error(err))
		}

		storage, err := pgadapter.New(logger, db)
		if err != nil {
			logger.Fatal("failed to setup postgresql directory database:", zap.Error(err))
		}

		httpClient := nlxhttp.NewHTTPClient(certificate)
		directoryService := directory.New(&directory.NewDirectoryArgs{
			Logger:                                logger,
			TermsOfServiceURL:                     serveOpts.TermsOfServiceURL,
			Repository:                            storage,
			Clock:                                 &clock{},
			HTTPClient:                            httpClient,
			GetOrganizationInformationFromRequest: common_tls.GetOrganizationInfoFromRequest,
		})

		httpServer := http.NewServer(db, certificate, logger)

		server, err := NewServer(&NewServerArgs{
			Logger:                       logger,
			Address:                      serveOpts.ListenAddress,
			AddressPlain:                 serveOpts.ListenAddressPlain,
			Certificate:                  certificate,
			DirectoryService:             directoryService,
			DirectoryRegistrationService: directoryService,
			DirectoryInspectionService:   directoryService,
			HTTPServer:                   httpServer,
		})
		if err != nil {
			logger.Fatal("could not start server", zap.Error(err))
		}

		p.Wait()

		logger.Info("starting graceful shutdown")

		gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err = server.Shutdown(gracefulCtx)
		if err != nil {
			logger.Error("could not shutdown server", zap.Error(err))
		}

		err = db.Close()
		if err != nil {
			logger.Error("could not shutdown db", zap.Error(err))
		}
	},
}
