// Copyright © VNG Realisatie 2023
// Licensed under the EUPL

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"go.nlx.io/nlx/common/cmd"
	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/process"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	grpcdirectory "go.nlx.io/nlx/directory-ui/adapters/directory/grpc"
	zaplogger "go.nlx.io/nlx/directory-ui/adapters/logger/zap"
	uiport "go.nlx.io/nlx/directory-ui/ports/ui"
	"go.nlx.io/nlx/directory-ui/service"
)

var serveOpts struct {
	ListenAddress    string
	DirectoryAddress string
	Environment      string
	StaticPath       string

	logoptions.LogOptions
	cmd.TLSOrgOptions
}

//nolint:gochecknoinits,funlen,gocyclo // this is the recommended way to use cobra, also a lot of flags..
func init() {
	serveCommand.Flags().StringVarP(&serveOpts.ListenAddress, "listen-address", "", "127.0.0.1:3001", "Address for the UI to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs.")
	serveCommand.Flags().StringVarP(&serveOpts.DirectoryAddress, "directory-address", "", "directory-api.shared.nlx.local:443", "URL to the Directory API")
	serveCommand.Flags().StringVarP(&serveOpts.Environment, "environment", "", "local", "Environment of this UI. local, acc, demo, preprod or prod.")
	serveCommand.Flags().StringVarP(&serveOpts.StaticPath, "static-path", "", "public", "Path to the static web files")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogType, "log-type", "", "live", "Set the logging config. See NewProduction and NewDevelopment at https://godoc.org/go.uber.org/zap#Logger.")
	serveCommand.Flags().StringVarP(&serveOpts.LogOptions.LogLevel, "log-level", "", "", "Set loglevel")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.NLXRootCert, "tls-nlx-root-cert", "", "", "Absolute or relative path to the NLX CA root cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgCertFile, "tls-org-cert", "", "", "Absolute or relative path to the Organization cert .pem")
	serveCommand.Flags().StringVarP(&serveOpts.TLSOrgOptions.OrgKeyFile, "tls-org-key", "", "", "Absolute or relative path to the Organization key .pem")
}

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start the UI",
	Run: func(cmd *cobra.Command, args []string) {
		p := process.NewProcess()

		logger, err := zaplogger.New(serveOpts.LogOptions.LogLevel, serveOpts.LogOptions.LogType)
		if err != nil {
			log.Fatalf("failed to create logger: %v", err)
		}

		logger.Info(fmt.Sprintf("version info: version: %s source-hash: %s", version.BuildVersion, version.BuildSourceHash))

		if errValidate := common_tls.VerifyPrivateKeyPermissions(serveOpts.OrgKeyFile); errValidate != nil {
			logger.Warn("invalid organization key permissions", err)
		}

		certificate, err := common_tls.NewBundleFromFiles(serveOpts.OrgCertFile, serveOpts.OrgKeyFile, serveOpts.NLXRootCert)
		if err != nil {
			logger.Fatal("loading certificate", err)
		}

		ctx := context.Background()

		directoryClient, err := grpcdirectory.NewClient(ctx, serveOpts.DirectoryAddress, certificate)
		if err != nil {
			logger.Fatal("create directory client", err)
		}

		directoryRepository := grpcdirectory.New(directoryClient)

		app, err := service.NewApplication(&service.NewApplicationArgs{
			Context:             ctx,
			DirectoryRepository: directoryRepository,
		})
		if err != nil {
			logger.Fatal("could not create application", err)
		}

		workDir, err := os.Getwd()
		if err != nil {
			logger.Fatal("failed to get work dir", err)
		}

		staticFilesPath := filepath.Join(workDir, serveOpts.StaticPath)

		uiServer, err := uiport.New(serveOpts.Environment, staticFilesPath, logger, app)

		go func() {
			err = uiServer.ListenAndServe(serveOpts.ListenAddress)
			if err != nil {
				logger.Fatal("could not listen and serve", err)
			}
		}()

		p.Wait()

		logger.Info("starting graceful shutdown")

		gracefulCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		err = uiServer.Shutdown(gracefulCtx)
		if err != nil {
			logger.Error("could not shutdown server", err)
		}

		err = directoryRepository.Shutdown()
		if err != nil {
			logger.Error("could not shutdown directory repository", err)
		}
	},
}
