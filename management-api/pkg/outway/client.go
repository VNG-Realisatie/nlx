// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package outway

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	outwayapi "go.nlx.io/nlx/outway/api"
)

const component = "nlx-management"

var (
	userAgent   = component + "/" + version.BuildVersion
	dialTimeout = 10 * time.Second
)

type Client interface {
	outwayapi.OutwayClient
}

type client struct {
	outwayapi.OutwayClient
}

func NewClient(ctx context.Context, outwayAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
		grpc.WithUnaryInterceptor(timeoutUnaryInterceptor),
	}

	ctx = nlxversion.NewGRPCContext(ctx, component)

	outwayConn, err := grpc.DialContext(ctx, outwayAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	c := &client{
		outwayapi.NewOutwayClient(outwayConn),
	}

	return c, nil
}

func timeoutUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, _ = context.WithTimeout(ctx, dialTimeout) // nolint:govet // cancel function is used by the invoker

	return invoker(ctx, method, req, reply, cc, opts...)
}
