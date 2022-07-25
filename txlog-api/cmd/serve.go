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
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	pgadapter "go.nlx.io/nlx/txlog-api/adapters/storage/postgres"
	"go.nlx.io/nlx/txlog-api/pkg/api"
)

var serveOpts struct {
	ListenAddress      string
	ListenAddressPlain string
	PostgresDSN        string

	logoptions.LogOptions
	cmd.TLSOptions
}

//nolint:gochecknoinits,funlen,gocyclo // this is the recommended way to use cobra, also a lot of flags..
func init() {
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddress, "listen-address", "", "127.0.0.1:8443", "Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddressPlain, "listen-address-plain", "", "127.0.0.1:8080", "Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogType, "log-type", "", "live", "Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogLevel, "log-level", "", "", "Override the default loglevel as set by --log-type.")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.RootCertFile, "tls-root-cert", "", "", "Absolute or relative path to the CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.CertFile, "tls-cert", "", "", "Absolute or relative path to the cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.KeyFile, "tls-key", "", "", "Absolute or relative path to the key .pem")
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start the API",
	Run: func(cmd *cobra.Command, args []string) {
		p := process.NewProcess()

		zapConfig := serveOpts.LogOptions.ZapConfig()
		logger, err := zapConfig.Build()
		if err != nil {
			log.Fatalf("failed to create new zap logger: %v", err)
		}

		logger.Info("version info", zap.String("version", version.BuildVersion), zap.String("source-hash", version.BuildSourceHash))
		logger = logger.With(zap.String("version", version.BuildVersion))

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.KeyFile); errValidate != nil {
			logger.Warn("invalid internal PKI key permissions", zap.Error(errValidate), zap.String("file-path", serveOpts.KeyFile))
		}

		certificate, err := common_tls.NewBundleFromFiles(serveOpts.CertFile, serveOpts.KeyFile, serveOpts.RootCertFile)
		if err != nil {
			logger.Fatal("loading TLS files", zap.Error(err))
		}

		db, err := pgadapter.NewPostgreSQLConnection(serveOpts.PostgresDSN)
		if err != nil {
			logger.Fatal("can not create db connection:", zap.Error(err))
		}

		storage, err := pgadapter.New(logger, db)
		if err != nil {
			logger.Fatal("failed to setup postgresql txlog database:", zap.Error(err))
		}

		server, err := api.NewAPI(logger, certificate, storage)
		if err != nil {
			logger.Fatal("could not start server", zap.Error(err))
		}

		go func() {
			err = server.ListenAndServe(serveOpts.ListenAddress, serveOpts.ListenAddressPlain)
			if err != nil {
				logger.Fatal("could not listen and serve", zap.Error(err))
			}
		}()

		p.Wait()

		logger.Info("starting graceful shutdown")

		gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err = server.Shutdown(gracefulCtx)
		if err != nil {
			logger.Error("could not shutdown server", zap.Error(err))
		}

		err = storage.Shutdown()
		if err != nil {
			logger.Error("could not shutdown storage", zap.Error(err))
		}
	},
}
