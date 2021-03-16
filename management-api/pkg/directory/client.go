// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/nlxversion"
	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/common/version"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const component = "nlx-management"

var (
	userAgent   = component + "/" + version.BuildVersion
	dialTimeout = 10 * time.Second
)

type Client interface {
	inspectionapi.DirectoryInspectionClient
	registrationapi.DirectoryRegistrationClient

	GetOrganizationInwayProxyAddress(ctx context.Context, organizationName string) (string, error)
}

type client struct {
	inspectionapi.DirectoryInspectionClient
	registrationapi.DirectoryRegistrationClient
}

func NewClient(ctx context.Context, inspectionAddress, registrationAddress string, cert *common_tls.CertificateBundle) (Client, error) {
	dialCredentials := credentials.NewTLS(cert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
		grpc.WithUserAgent(userAgent),
		grpc.WithUnaryInterceptor(timeoutUnaryInterceptor),
	}

	ctx = nlxversion.NewGRPCContext(ctx, component)

	inspectionConn, err := grpc.DialContext(ctx, inspectionAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	registrationConn, err := grpc.DialContext(ctx, registrationAddress, dialOptions...)
	if err != nil {
		return nil, err
	}

	c := &client{
		inspectionapi.NewDirectoryInspectionClient(inspectionConn),
		registrationapi.NewDirectoryRegistrationClient(registrationConn),
	}

	return c, nil
}

func (c *client) GetOrganizationInwayProxyAddress(ctx context.Context, organizationName string) (string, error) {
	response, err := c.GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
		OrganizationName: organizationName,
	})
	if err != nil {
		return "", err
	}

	return computeInwayProxyAddress(response.Address)
}

func computeInwayProxyAddress(address string) (string, error) {
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
