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
	testPublicKeyDER = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxahK/ruBG74MZ2/Z71llWS1OMJy6xs9qpQY7YC7C+u59JoqNdoWToV6EZRfPYzh61BWsyKlqvkl11HhD6HVkWmdidYJtmoJqmFeWm02a6RkP4XbBiOCm9xxX/xlZRWubTCswaLkfmlI2IYxgpLIPQuvO2nbor8YG4dS7u7u7yrtl1dOBD1utlMCzX5j8vG+BaHUqE1kBWIE5kg9ogeVfwa/w30EPcD+5gdknn5uGoTFP/xi6WiZ+6MJli1CPjrHX0N73ZMSdgHK+4jk8KdrzFou5sNtCl+CTdzhDhwYJxJv/McsgqPfXsOdk0T3QUcCqWsawJ8VblJYYwyj1WW7lbJSygJjvOTG+C2+vbht3mKvimKpx/+8S/Zg+g7nen//SvFQhe2wI7Eaottgk/abU6i3ntvSty4EyxFPnchKa7EXeFAsp4stO0Q5iTE4rEdDotwaWrmcN54UQr2ZOVPJ/BGGG6SxeciX9jB9I1FHBngMyiXVDgMlgGa9Ke3y1V+Yaqh3LOp6JXnjXp50Ke0ncCMa7tBd6GGJqV4hl3daYj7yyBWzB3E2d/u+gJx1e9mxqgA0V7nidh2CRelHtczhCO5/DpYFGnjKm4YMkzSb7CxRDrL2OJeyvM3tKyRZES5eEiedMcpjvm5ULzZeCp2r3P72Jy9qTigqNYoIHBYMpFzUCAwEAAQ=="
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
	assert.Equal(t, []string{"g+jpuLAMFzM09tOZpb0Ehslhje4S/IsIxSWsS4E16Yc="}, md.Get("nlx-public-key-fingerprint"))
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
