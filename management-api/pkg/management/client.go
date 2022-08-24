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
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func NewClient(ctx context.Context, inwayAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
	}

	var grpcTimeout = 5 * time.Second

	timeoutCtx, cancel := context.WithTimeout(ctx, grpcTimeout)

	grpcCtx := nlxversion.NewGRPCContext(timeoutCtx, component)

	conn, err := grpc.DialContext(grpcCtx, inwayAddress, dialOptions...)
	if err != nil {
		cancel()
		return nil, err
	}

	c := client{
		conn:                       conn,
		cancel:                     cancel,
		AccessRequestServiceClient: external.NewAccessRequestServiceClient(conn),
		DelegationServiceClient:    external.NewDelegationServiceClient(conn),
	}

	return &c, nil
}

func (client *client) Close() error {
	client.cancel()
	return client.conn.Close()
}
