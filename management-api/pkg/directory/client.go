// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
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
	userAgent   = component + "/" + version.BuildVersion
	dialTimeout = 10 * time.Second
)

type Client interface {
	directoryapi.DirectoryClient

	GetOrganizationInwayProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error)
}

type client struct {
	directoryapi.DirectoryClient
}

func NewClient(ctx context.Context, directoryAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
		grpc.WithUnaryInterceptor(timeoutUnaryInterceptor),
	}

	ctx = nlxversion.NewGRPCContext(ctx, component)

	directoryConn, err := grpc.DialContext(ctx, directoryAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	c := &client{
		directoryapi.NewDirectoryClient(directoryConn),
	}

	return c, nil
}

func (c *client) GetOrganizationInwayProxyAddress(ctx context.Context, organizationSerialNumber string) (string, error) {
	response, err := c.GetOrganizationInway(ctx, &directoryapi.GetOrganizationInwayRequest{
		OrganizationSerialNumber: organizationSerialNumber,
	})
	if err != nil {
		return "", err
	}

	return computeInwayProxyAddress(response.Address)
}

func computeInwayProxyAddress(address string) (string, error) {
	if address == "" {
		return "", errors.New("empty inway address provided")
	}

	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address: %w", err)
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address port: %w", err)
	}

	return fmt.Sprintf("%s:%d", host, portNum+1), nil
}

func timeoutUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, _ = context.WithTimeout(ctx, dialTimeout) // nolint:govet // cancel function is used by the invoker

	return invoker(ctx, method, req, reply, cc, opts...)
}
