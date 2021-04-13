// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package grpcproxy

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/tls"
)

type Proxy struct {
	methods *methodSet
	server  *grpc.Server
	conn    *grpc.ClientConn
	logger  *zap.Logger
}

type methodSet map[string]bool

func (m methodSet) set(service, method string) {
	k := fullMethod(service, method)
	m[k] = true
}

func (m methodSet) has(fullMethod string) bool {
	_, ok := m[fullMethod]
	return ok
}

func fullMethod(service, method string) string {
	return "/" + service + "/" + method
}

func New(ctx context.Context, logger *zap.Logger, upstreamAddress string, serverCert, clientCert *tls.CertificateBundle, dialOpts ...grpc.DialOption) (*Proxy, error) {
	codec := &passthroughCodec{}
	dialOpts = append(
		dialOpts,
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)),
		grpc.WithTransportCredentials(credentials.NewTLS(clientCert.TLSConfig())),
	)

	cc, err := grpc.DialContext(ctx, upstreamAddress, dialOpts...)
	if err != nil {
		return nil, fmt.Errorf("connect to %v: %w", upstreamAddress, err)
	}

	p := &Proxy{
		methods: &methodSet{},
		conn:    cc,
		logger:  logger,
	}

	tlsConfig := serverCert.TLSConfig(serverCert.WithTLSClientAuth())

	p.server = grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.CustomCodec(codec),
		grpc.StreamInterceptor(p.streamInterceptor),
		grpc.UnknownServiceHandler(p.serviceHandler),
	)

	return p, nil
}

func (p *Proxy) RegisterService(s *grpc.ServiceDesc) {
	for _, m := range s.Methods {
		p.methods.set(s.ServiceName, m.MethodName)
	}
}

func (p *Proxy) Serve(l net.Listener) error {
	return p.server.Serve(l)
}

func (p *Proxy) Stop() {
	p.server.GracefulStop()
}

func (p *Proxy) streamInterceptor(srv interface{}, serverStream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if !p.methods.has(info.FullMethod) {
		p.logger.Debug("unknown service/method", zap.String("fullMethod", info.FullMethod))
		return status.Error(codes.Unimplemented, fmt.Sprintf("unknown service/method %v", info.FullMethod))
	}

	ctx := serverStream.Context()

	pr, ok := peer.FromContext(ctx)
	if !ok {
		p.logger.Warn("peer is missing from context")
		return status.Error(codes.Internal, "peer is missing from context")
	}

	tlsInfo, ok := pr.AuthInfo.(credentials.TLSInfo)
	if !ok || tlsInfo.State.PeerCertificates == nil {
		p.logger.Warn("tls peer certificate is missing")
		return status.Error(codes.Internal, "tls peer certificate is missing")
	}

	peerCert := tlsInfo.State.PeerCertificates[0]

	if len(peerCert.Subject.Organization) == 0 {
		p.logger.Warn("certificate is missing organization")
		return status.Error(codes.Unauthenticated, "certificate is missing organization")
	}

	publicKey, ok := peerCert.PublicKey.(*rsa.PublicKey)
	if !ok {
		p.logger.Warn("invalid format for public key")
		return status.Error(codes.Internal, "invalid format for public key")
	}

	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		p.logger.Warn("invalid format for public key")
		return status.Error(codes.Internal, "invalid format for public key")
	}

	streamInfo := &streamInfo{
		fullMethod:           info.FullMethod,
		peerAddr:             pr.Addr.String(),
		organizationName:     peerCert.Subject.Organization[0],
		publicKeyDER:         base64.StdEncoding.EncodeToString(publicKeyDER),
		publicKeyFingerprint: tls.PublicKeyFingerprint(peerCert),
	}

	w := &wrappedServerStream{serverStream, setStreamInfo(ctx, streamInfo)}

	p.logger.Debug("handeling method", zap.String("method", info.FullMethod))

	return handler(srv, w)
}

// serviceHandler is gRPC StreamHandler handling incomming server streams.
func (p *Proxy) serviceHandler(srv interface{}, serverStream grpc.ServerStream) error {
	ctx := serverStream.Context()
	info := extractStreamInfo(ctx)

	if info == nil {
		p.logger.Error("stream info is missing")
		return status.Error(codes.Internal, "stream info is missing")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		p.logger.Error("metadata missing")
		return status.Error(codes.Internal, "metadata missing")
	}

	forwarded := fmt.Sprintf("for=%s", info.peerAddr)
	if authority, ok := md[":authority"]; ok {
		forwarded = fmt.Sprintf("%s,host=%s", forwarded, authority[0])
	}

	clientMD := md.Copy()
	clientMD.Set("forwarded", forwarded)
	clientMD.Set("nlx-organization", info.organizationName)
	clientMD.Set("nlx-public-key-der", info.publicKeyDER)
	clientMD.Set("nlx-public-key-fingerprint", info.publicKeyFingerprint)

	clientCtx := metadata.NewOutgoingContext(ctx, clientMD)
	clientCtx, cancelClient := context.WithCancel(clientCtx)

	defer cancelClient()

	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDesc, p.conn, info.fullMethod)
	if err != nil {
		p.logger.Error("unable to create client stream", zap.Error(err))

		return status.Error(codes.Unavailable, "")
	}

	return p.proxyStreamToUpstream(serverStream, clientStream)
}
