// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	directoryapi "go.nlx.io/nlx/directory-api/api"
)

const component = "nlx-management"

var (
	userAgent = component + "/" + version.BuildVersion
)

type Client interface {
	directoryapi.DirectoryClient

	GetOrganizationInwayProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error)
	Close() error
}

type client struct {
	directoryapi.DirectoryClient
	conn   *grpc.ClientConn
	cancel context.CancelFunc
}

func NewClient(ctx context.Context, directoryAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
	}

	var grpcTimeout = 5 * time.Second

	timeoutCtx, cancel := context.WithTimeout(ctx, grpcTimeout)

	grpcCtx := nlxversion.NewGRPCContext(timeoutCtx, component)

	directoryConn, err := grpc.DialContext(grpcCtx, directoryAddress, dialOptions...)
	if err != nil {
		cancel()
		return nil, err
	}

	c := &client{
		DirectoryClient: directoryapi.NewDirectoryClient(directoryConn),
		cancel:          cancel,
		conn:            directoryConn,
	}

	return c, nil
}

func (c *client) GetOrganizationInwayProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	response, err := c.GetOrganizationManagementAPIProxyAddress(ctx, &directoryapi.GetOrganizationManagementAPIProxyAddressRequest{
		OrganizationSerialNumber: organizationSerialNumber,
	})
	if err != nil {
		return "", err
	}

	return response.Address, nil
}

func (c *client) Close() error {
	c.cancel()

	return c.conn.Close()
}
