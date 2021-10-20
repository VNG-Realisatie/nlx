// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy_test

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/inway/grpcproxy/test"
	common_testing "go.nlx.io/nlx/testing/testingutils"
)

func TestUnknownServiceMethod(t *testing.T) {
	// nolint:dogsled // we only need the client
	_, _, c, _ := setup(t, nil)

	ctx := context.Background()
	resp, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})

	assert.Nil(t, resp)
	assert.Error(t, err)

	st := status.Convert(err)
	assert.Equal(t, codes.Unimplemented, st.Code())
	assert.Equal(t, "unknown service/method /grpcproxy.test.TestService/Test", st.Message())
}

func TestRegisteredService(t *testing.T) {
	p, s, c, _ := setup(t, nil)

	p.RegisterService(test.GetTestServiceDesc())

	ctx := context.Background()

	resp, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})
	assert.NoError(t, err)
	assert.Equal(t, "Foo", resp.Name)

	resp, err = c.Test(ctx, &test.TestRequest{Name: "Bar"})
	assert.NoError(t, err)
	assert.Equal(t, "Bar", resp.Name)

	assert.Len(t, s.svc.reqs, 2)
}

func TestMetadataToUpstreamCantBeOverridden(t *testing.T) {
	p, s, c, certBundle := setup(t, nil)
	p.RegisterService(test.GetTestServiceDesc())

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"forwarded", "spoofed-forwarded",
		"forwarded", "spoofed-forwarded",
		"nlx-organization-name", "spoofed-organization",
		"nlx-organization-name", "spoofed-organization",
		"nlx-public-key-der", "spoofed-public-key",
		"nlx-public-key-der", "spoofed-public-key",
		"nlx-public-key-fingerprint", "spoofed-fingerprint",
		"nlx-public-key-fingerprint", "spoofed-fingerprint",
		"some-key", "foo",
		"some-other-key", "bar",
		"some-other-key", "foobar",
	)

	_, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})
	assert.NoError(t, err)

	md := s.svc.reqs[0].md

	publicKeyDER, err := x509.MarshalPKIXPublicKey(certBundle.PublicKey())
	assert.NoError(t, err)

	publicKeyDEREncoded := base64.StdEncoding.EncodeToString(publicKeyDER)

	assert.Equal(t, []string{"for=bufconn,host=inway.test"}, md.Get("forwarded"))
	assert.Equal(t, []string{"nlx-test"}, md.Get("nlx-organization-name"))
	assert.Equal(t, []string{publicKeyDEREncoded}, md.Get("nlx-public-key-der"))
	assert.Equal(t, []string{certBundle.PublicKeyFingerprint()}, md.Get("nlx-public-key-fingerprint"))
	assert.Equal(t, []string{"foo"}, md.Get("some-key"))
	assert.Equal(t, []string{"bar", "foobar"}, md.Get("some-other-key"))
}

func TestMissingOrganization(t *testing.T) {
	pkiDir := filepath.Join("..", "..", "testing", "pki")

	clientCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutName)
	assert.NoError(t, err)

	p, s, c, _ := setup(t, clientCert)
	p.RegisterService(test.GetTestServiceDesc())

	ctx := context.Background()
	resp, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})

	assert.Nil(t, resp)
	assert.Error(t, err)

	st := status.Convert(err)
	assert.Equal(t, codes.Unauthenticated, st.Code())
	assert.Equal(t, "certificate is missing organization", st.Message())

	assert.Empty(t, s.svc.reqs)
}

// nolint:dupl // testing different property
func TestMissingSerialNumber(t *testing.T) {
	pkiDir := filepath.Join("..", "..", "testing", "pki")

	clientCert, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgWithoutSerialNumber)
	assert.NoError(t, err)

	p, s, c, _ := setup(t, clientCert)
	p.RegisterService(test.GetTestServiceDesc())

	ctx := context.Background()
	resp, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})

	assert.Nil(t, resp)
	assert.Error(t, err)

	st := status.Convert(err)
	assert.Equal(t, codes.Unauthenticated, st.Code())
	assert.Equal(t, "certificate is missing serial number", st.Message())

	assert.Empty(t, s.svc.reqs)
}

func setup(t *testing.T, clientCert *tls.CertificateBundle) (*grpcproxy.Proxy, *testServer, test.TestServiceClient, *tls.CertificateBundle) {
	ctx := context.Background()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	pkiDir := filepath.Join("..", "..", "testing", "pki")

	certBundle, err := common_testing.GetCertificateBundle(pkiDir, common_testing.OrgNLXTest)
	assert.NoError(t, err)

	s := newTestServer(t, certBundle)
	s.start()

	p, err := grpcproxy.New(
		ctx,
		logger,
		s.address(),
		certBundle,
		certBundle,
		grpc.WithContextDialer(s.dialer),
	)
	assert.NoError(t, err)

	l := bufconn.Listen(bufferSize)

	go func() {
		log.Println(p.Serve(l))
	}()

	if clientCert == nil {
		clientCert = certBundle
	}

	c := newTestClient(t, l, clientCert)

	return p, s, c, certBundle
}
