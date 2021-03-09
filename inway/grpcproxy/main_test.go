package grpcproxy_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"

	"go.nlx.io/nlx/common/tls"

	"go.nlx.io/nlx/inway/grpcproxy/test"
)

const bufferSize = 1024 * 1024

type rpcRequest struct {
	req *test.TestRequest
	md  metadata.MD
}

type rpcResponse struct {
	resp *test.TestResponse
	err  error
}

type testService struct {
	test.UnimplementedTestServiceServer

	reqs []*rpcRequest
	resp *rpcResponse
}

// Test implements the test.TestServiceServer interface
func (s *testService) Test(ctx context.Context, req *test.TestRequest) (*test.TestResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	s.reqs = append(s.reqs, &rpcRequest{req, md})

	if s.resp != nil {
		return s.resp.resp, s.resp.err
	}

	return &test.TestResponse{Name: req.Name}, nil
}

type testServer struct {
	srv *grpc.Server
	l   *bufconn.Listener
	svc *testService
}

func (t *testServer) start() {
	go func() {
		if err := t.srv.Serve(t.l); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
func (t *testServer) stop() {
	t.srv.GracefulStop()
}

func (t *testServer) address() string {
	return "inway.test"
}

func (t *testServer) dialer(context.Context, string) (net.Conn, error) {
	return t.l.Dial()
}

func newTestServer(t *testing.T, cert *tls.CertificateBundle) *testServer {
	tlsConfig := cert.TLSConfig(cert.WithTLSClientAuth())
	srv := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
	)

	s := &testServer{
		srv: srv,
		l:   bufconn.Listen(bufferSize),
		svc: &testService{},
	}

	test.RegisterTestServiceServer(s.srv, s.svc)

	t.Cleanup(func() {
		s.stop()
	})

	return s
}

type testClient struct {
	test.TestServiceClient
	cc *grpc.ClientConn
}

func newTestClient(t *testing.T, l *bufconn.Listener, cert *tls.CertificateBundle) *testClient {
	dialer := func(context.Context, string) (net.Conn, error) {
		return l.Dial()
	}

	ctx := context.Background()

	cc, err := grpc.DialContext(ctx, "inway.test", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(credentials.NewTLS(cert.TLSConfig())))
	assert.NoError(t, err)

	c := &testClient{
		test.NewTestServiceClient(cc),
		cc,
	}

	t.Cleanup(func() {
		c.cc.Close()
	})

	return c
}
