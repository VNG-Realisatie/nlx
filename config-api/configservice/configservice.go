package configservice

import (
	"context"
	"fmt"
	"net"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

// ConfigService handles all requests for the config api
type ConfigService struct {
	connections          map[string]*component
	organization         string
	logger               *zap.Logger
	etcdCli              *clientv3.Client
	transportCredentials nlxTransportCredentials
}

type component struct {
	name       string
	kind       string
	connection *connection
}

// New creates new ConfigService
func New(logger *zap.Logger, etcdClient *clientv3.Client, organization string, transportCredentials credentials.TransportCredentials) *ConfigService {
	return &ConfigService{
		etcdCli:              etcdClient,
		connections:          make(map[string]*component),
		organization:         organization,
		transportCredentials: newTransportCredentials(transportCredentials, organization),
		logger:               logger,
	}
}

// GetTransportCredentials returns the transportcredentials to be used by GRPC
func (c ConfigService) GetTransportCredentials() credentials.TransportCredentials {
	return c.transportCredentials
}

// AuthorizationStreamInterceptor intercepts streams and verifies the organization
func (c ConfigService) AuthorizationStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	peer, ok := peer.FromContext(ss.Context())
	if ok {
		tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
		if tlsInfo.State.PeerCertificates[0].Subject.Organization[0] != c.organization {
			return fmt.Errorf("organization does not match")
		}
	}

	return handler(srv, ss)
}

type nlxTransportCredentials struct {
	organization string
	creds        credentials.TransportCredentials
}

func (t nlxTransportCredentials) ClientHandshake(c context.Context, s string, con net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return t.ClientHandshake(c, s, con)
}

func (t nlxTransportCredentials) Info() credentials.ProtocolInfo {
	return t.creds.Info()
}
func (t nlxTransportCredentials) Clone() credentials.TransportCredentials {
	return t.creds.Clone()
}
func (t nlxTransportCredentials) OverrideServerName(n string) error {
	return t.creds.OverrideServerName(n)
}

func (t nlxTransportCredentials) ServerHandshake(con net.Conn) (net.Conn, credentials.AuthInfo, error) {
	con, authInfo, err := t.creds.ServerHandshake(con)
	if err != nil {
		return con, authInfo, err
	}
	tlsInfo := authInfo.(credentials.TLSInfo)
	if tlsInfo.State.PeerCertificates[0].Subject.Organization[0] != t.organization {
		return nil, nil, fmt.Errorf("organization does not match")
	}
	return con, authInfo, err

}

func newTransportCredentials(c credentials.TransportCredentials, organization string) nlxTransportCredentials {
	return nlxTransportCredentials{
		creds:        c,
		organization: organization,
	}
}
