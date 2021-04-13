// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy_test

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"go.nlx.io/nlx/common/tls"

	"go.nlx.io/nlx/inway/grpcproxy"
	"go.nlx.io/nlx/inway/grpcproxy/test"
)

var (
	pkiDir           = filepath.Join("..", "..", "testing", "pki")
	testPublicKeyDER = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArN5xGkM73tJsCpKny59e5lXNRY+eT0sbWyEGsR1qIPRKmLSiRHl3xMsovn5mo6jN3eeK/Q4wKd6Ae5XGzP63pTG6U5KVVB74eQxSFfV3UEOrDaJ78X5mBZO+Ku21V2QFr44tvMh5IZDX3RbMB/4Kad6sapmSF00HWrqTVMkrEsZ98DTb5nwGLh3kISnct4tLyVSpsl9s1rtkSgGUcs1TIvWxS2D2mOsSL1HRdUNcFQmzchbfG87kXPvicoOISAZDJKDqWp3iuH0gJpQ+XMBfmcD90I7Z/cRQjWP3P93B3V06cJkd00cEIRcIQqF8N+lE01H88Fi+wePhZRy92NP54wIDAQAB"
)

func setup(t *testing.T, clientCert *tls.CertificateBundle) (*grpcproxy.Proxy, *testServer, test.TestServiceClient) {
	ctx := context.Background()
	logger, _ := zap.NewDevelopment()

	cert, err := tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-nlx-test-chain.pem"),
		filepath.Join(pkiDir, "org-nlx-test-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
	assert.NoError(t, err)

	s := newTestServer(t, cert)
	s.start()

	p, err := grpcproxy.New(
		ctx,
		logger,
		s.address(),
		cert,
		cert,
		grpc.WithContextDialer(s.dialer),
	)
	assert.NoError(t, err)

	l := bufconn.Listen(bufferSize)

	go func() {
		log.Println(p.Serve(l))
	}()

	if clientCert == nil {
		clientCert = cert
	}

	c := newTestClient(t, l, clientCert)

	return p, s, c
}

func TestUnknownServiceMethod(t *testing.T) {
	_, _, c := setup(t, nil)

	ctx := context.Background()
	resp, err := c.Test(ctx, &test.TestRequest{Name: "Foo"})

	assert.Nil(t, resp)
	assert.Error(t, err)

	st := status.Convert(err)
	assert.Equal(t, codes.Unimplemented, st.Code())
	assert.Equal(t, "unknown service/method /grpcproxy.test.TestService/Test", st.Message())
}

func TestRegisteredService(t *testing.T) {
	p, s, c := setup(t, nil)

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

// TestMetadata tests that metadata send to upstream can't be overridden
func TestMetadata(t *testing.T) {
	p, s, c := setup(t, nil)
	p.RegisterService(test.GetTestServiceDesc())

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"forwarded", "spoofed-forwarded",
		"forwarded", "spoofed-forwarded",
		"nlx-organization", "spoofed-organization",
		"nlx-organization", "spoofed-organization",
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

	assert.Equal(t, []string{"for=bufconn,host=inway.test"}, md.Get("forwarded"))
	assert.Equal(t, []string{"nlx-test"}, md.Get("nlx-organization"))
	assert.Equal(t, []string{testPublicKeyDER}, md.Get("nlx-public-key-der"))
	assert.Equal(t, []string{"60igp6kiaIF14bQCdNiPPhiP3XJ95qLFhAFI1emJcm4="}, md.Get("nlx-public-key-fingerprint"))
	assert.Equal(t, []string{"foo"}, md.Get("some-key"))
	assert.Equal(t, []string{"bar", "foobar"}, md.Get("some-other-key"))
}

func TestMissingOrganization(t *testing.T) {
	clientCert, err := tls.NewBundleFromFiles(
		filepath.Join(pkiDir, "org-without-name-chain.pem"),
		filepath.Join(pkiDir, "org-without-name-key.pem"),
		filepath.Join(pkiDir, "ca-root.pem"),
	)
	assert.NoError(t, err)

	p, s, c := setup(t, clientCert)
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
