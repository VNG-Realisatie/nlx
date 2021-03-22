// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package management

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/management-api/api/external"
)

const component = "nlx-management"

var (
	userAgent   = component + "/" + version.BuildVersion
	dialTimeout = 10 * time.Second
)

type Client interface {
	external.AccessRequestServiceClient
	external.DelegationServiceClient
	Close() error
}

type client struct {
	external.AccessRequestServiceClient
	external.DelegationServiceClient
	conn *grpc.ClientConn
}

func NewClient(ctx context.Context, inwayAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
		grpc.WithUnaryInterceptor(timeoutUnaryInterceptor),
	}

	ctx = nlxversion.NewGRPCContext(ctx, component)

	conn, err := grpc.DialContext(ctx, inwayAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	c := client{
		conn: conn,
	}

	return &c, nil
}

func (client *client) Close() error {
	return client.conn.Close()
}

func timeoutUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, _ = context.WithTimeout(ctx, dialTimeout) // nolint:govet // cancel function is used by the invoker

	return invoker(ctx, method, req, reply, cc, opts...)
}
