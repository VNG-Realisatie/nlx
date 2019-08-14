package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	flags "github.com/svent/go-flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tlsconfig"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
)

var options struct {
	ListenAddress        string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:8443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	EtcdConnectionString string `long:"etcd-connection-string" env:"ETCD_CONNECTION_STRING" description:"A comma separated list of etcd backends." required:"true"`
	orgtls.TLSOptions
	logoptions.LogOptions
}

func main() {
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

	config := options.LogOptions.ZapConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new zap logger: %v", err)
	}

	p := process.NewProcess(logger)

	db, err := configservice.NewEtcdConfigDatabase(logger, p, strings.Split(options.EtcdConnectionString, ","))
	if err != nil {
		logger.Fatal("failed to setup database", zap.Error(err))
	}

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)

	certKeyPair, err := tls.LoadX509KeyPair(options.TLSOptions.OrgCertFile, options.TLSOptions.OrgKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair", zap.Error(err))
	}

	certPool, _, err := orgtls.Load(options.TLSOptions)
	if err != nil {
		logger.Fatal("failed to load certifcates", zap.Error(err))
	}

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return status.Error(codes.Internal, fmt.Sprintf("%s", p))
		}),
	}

	// prepare grpc server options
	serverTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{certKeyPair},
		ClientCAs:    certPool,
		NextProtos:   []string{"h2"},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsconfig.ApplyDefaults(serverTLSConfig)

	transportCredentials := credentials.NewTLS(serverTLSConfig)
	confServer := configservice.New(logger, p, db)
	opts := []grpc.ServerOption{
		grpc.Creds(transportCredentials),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(recoveryOptions...),
		),
	}

	grpcServer := grpc.NewServer(
		opts...)

	configapi.RegisterConfigApiServer(grpcServer, confServer)

	startServer(p, grpcServer)
}

func startServer(p *process.Process, grpcServer *grpc.Server) {
	listen, err := net.Listen("tcp", options.ListenAddress)
	if err != nil {
		log.Fatal("failed to create listener", zap.Error(err))
	}
	p.CloseGracefully(func() error {
		grpcServer.GracefulStop()
		return nil
	})
	p.CloseGracefully(listen.Close)
	if err := grpcServer.Serve(listen); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal("error serving", zap.Error(err))
		}
	}
}
