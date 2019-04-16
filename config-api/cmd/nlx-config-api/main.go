package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	flags "github.com/svent/go-flags"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/logoptions"
	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/common/tlsconfig"
	"go.nlx.io/nlx/config-api/configapi"
	"go.nlx.io/nlx/config-api/configservice"
)

var options struct {
	ListenAddress        string `long:"listen-address" env:"LISTEN_ADDRESS" default:"0.0.0.0:443" description:"Address for the directory to listen on. Read https://golang.org/pkg/net/#Dial for possible tcp address specs."`
	ETCDConnectionString string `long:"etcd-connection-string" env:"ETCD-CONNECTION-STRING" description:"Connection string required to connect to the etcd config storage DB." required:"true"`
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
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{options.ETCDConnectionString},
		DialTimeout: time.Second,
	})
	if err != nil {
		logger.Fatal("failed to setup ETCD", zap.Error(err))
	}
	p.CloseGracefully(cli.Close)

	// setup zap connection for global grpc logging
	grpc_zap.ReplaceGrpcLogger(logger)

	certKeyPair, err := tls.LoadX509KeyPair(options.TLSOptions.OrgCertFile, options.TLSOptions.OrgKeyFile)
	if err != nil {
		logger.Fatal("failed to load x509 keypair", zap.Error(err))
	}

	certPool, cert, err := orgtls.Load(options.TLSOptions)
	if err != nil {
		logger.Fatal("failed to load certifcates", zap.Error(err))
	}

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logger.Warn("recovered from a panic in a grpc request handler", zap.ByteString("stack", debug.Stack()))
			return grpc.Errorf(codes.Internal, "%s", p)
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

	confServer := configservice.New(logger, cli, cert.Subject.Organization[0], credentials.NewTLS(serverTLSConfig))
	opts := []grpc.ServerOption{
		grpc.Creds(confServer.GetTransportCredentials()),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(recoveryOptions...),
			confServer.AuthorizationStreamInterceptor,
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
