// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package txlog

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	txlogapi "go.nlx.io/nlx/txlog-api/api"
)

const component = "nlx-management"

var (
	userAgent   = component + "/" + version.BuildVersion
	dialTimeout = 10 * time.Second
)

type Client interface {
	txlogapi.TXLogClient
	Enabled() bool
}

type client struct {
	txlogapi.TXLogClient
	enabled bool
}

func NewClient(ctx context.Context, txlogAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
		grpc.WithUnaryInterceptor(timeoutUnaryInterceptor),
	}

	ctx = nlxversion.NewGRPCContext(ctx, component)

	txlogConn, err := grpc.DialContext(ctx, txlogAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	c := &client{
		txlogapi.NewTXLogClient(txlogConn),
		true,
	}

	return c, nil
}

func (tx *client) Enabled() bool {
	return tx.enabled
}

func timeoutUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, _ = context.WithTimeout(ctx, dialTimeout) // nolint:govet // cancel function is used by the invoker

	return invoker(ctx, method, req, reply, cc, opts...)
}
