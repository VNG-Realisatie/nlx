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
	userAgent = component + "/" + version.BuildVersion
)

type Client interface {
	outwayapi.OutwayClient
	Close() error
}

type client struct {
	outwayapi.OutwayClient
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func NewClient(ctx context.Context, outwayAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
	}

	var grpcTimeout = 10 * time.Second

	timeoutCtx, cancel := context.WithTimeout(ctx, grpcTimeout)

	grpcCtx := nlxversion.NewGRPCContext(timeoutCtx, component)

	outwayConn, err := grpc.DialContext(grpcCtx, outwayAddress, dialOptions...)
	if err != nil {
		cancel()
		return nil, err
	}

	c := &client{
		OutwayClient: outwayapi.NewOutwayClient(outwayConn),
		conn:         outwayConn,
		cancel:       cancel,
	}

	return c, nil
}

func (c *client) Close() error {
	c.cancel()
	return c.conn.Close()
}
