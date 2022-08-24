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
	userAgent = component + "/" + version.BuildVersion
)

type Client interface {
	txlogapi.TXLogClient
	Close() error
}

type client struct {
	txlogapi.TXLogClient
	cancel context.CancelFunc
	conn   *grpc.ClientConn
}

func NewClient(ctx context.Context, txlogAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
	}

	var grpcTimeout = 10 * time.Second

	timeoutCtx, cancel := context.WithTimeout(ctx, grpcTimeout)

	grpcCtx := nlxversion.NewGRPCContext(timeoutCtx, component)

	txlogConn, err := grpc.DialContext(grpcCtx, txlogAddress, dialOptions...)
	if err != nil {
		cancel()
		return nil, err
	}

	c := &client{
		TXLogClient: txlogapi.NewTXLogClient(txlogConn),
		cancel:      cancel,
		conn:        txlogConn,
	}

	return c, nil
}

func (c *client) Close() error {
	c.cancel()

	return c.conn.Close()
}
