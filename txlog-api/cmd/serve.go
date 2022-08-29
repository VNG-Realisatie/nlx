// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/common/cmd"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	zaplogger "go.nlx.io/nlx/txlog-api/adapters/logger/zap"
	pgadapter "go.nlx.io/nlx/txlog-api/adapters/storage/postgres"
	ports_grpc "go.nlx.io/nlx/txlog-api/ports/grpc"
	"go.nlx.io/nlx/txlog-api/service"
)

var serveOpts struct {
	ListenAddress      string
	ListenAddressPlain string
	PostgresDSN        string

	logoptions.LogOptions
	cmd.TLSOptions
}

type clock struct{}

func (c *clock) Now() time.Time {
	return time.Now()
}

//nolint:gochecknoinits,funlen,gocyclo // this is the recommended way to use cobra, also a lot of flags..
func init() {
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddress, "listen-address", "", "127.0.0.1:8443", "Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddressPlain, "listen-address-plain", "", "127.0.0.1:8080", "Address for the api to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.PostgresDSN, "postgres-dsn", "", "", "Postgres Connection URL")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogType, "log-type", "", "live", "Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogLevel, "log-level", "", "", "Set loglevel")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.RootCertFile, "tls-root-cert", "", "", "Absolute or relative path to the CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.CertFile, "tls-cert", "", "", "Absolute or relative path to the cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOptions.KeyFile, "tls-key", "", "", "Absolute or relative path to the key .pem")
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start the API",
	Run: func(cmd *cobra.Command, args []string) {
		p := process.NewProcess()

		logger, err := zaplogger.New(serveOpts.LogOptions.LogLevel, serveOpts.LogOptions.LogType)
		if err != nil {
			log.Fatalf("failed to create logger: %v", err)
		}

		logger.Info(fmt.Sprintf("version info: version: %s source-hash: %s", version.BuildVersion, version.BuildSourceHash))

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.KeyFile); errValidate != nil {
			logger.Warn(fmt.Sprintf("invalid internal PKI key permissions: file-path: %s", serveOpts.KeyFile), err)
		}

		certificate, err := common_tls.NewBundleFromFiles(serveOpts.CertFile, serveOpts.KeyFile, serveOpts.RootCertFile)
		if err != nil {
			logger.Fatal("loading TLS files", err)
		}

		db, err := pgadapter.NewPostgreSQLConnection(serveOpts.PostgresDSN)
		if err != nil {
			logger.Fatal("can not create db connection:", err)
		}

		storage, err := pgadapter.New(db)
		if err != nil {
			logger.Fatal("failed to setup postgresql txlog database", err)
		}

		ctx := context.Background()

		app, err := service.NewApplication(&service.NewApplicationArgs{
			Context:    ctx,
			Logger:     logger,
			Repository: storage,
			Clock:      &clock{},
		})
		if err != nil {
			logger.Fatal("could not create application", err)
		}

		grpcServer := ports_grpc.New(logger, app, certificate)

		go func() {
			err = grpcServer.ListenAndServe(serveOpts.ListenAddress, serveOpts.ListenAddressPlain)
			if err != nil {
				logger.Fatal("could not listen and serve", err)
			}
		}()

		p.Wait()

		logger.Info("starting graceful shutdown")

		gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err = grpcServer.Shutdown(gracefulCtx)
		if err != nil {
			logger.Error("could not shutdown server", err)
		}

		err = storage.Shutdown()
		if err != nil {
			logger.Error("could not shutdown storage", err)
		}
	},
}
