// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/oidc"
)

// API handles incoming requests, authenticates them and forwards them to the config-api and txlog-api
type API struct {
	logger           *zap.Logger
	roots            *x509.CertPool
	orgCertKeyPair   *tls.Certificate
	process          *process.Process
	mux              *runtime.ServeMux
	configAPIAddress string
	authenticator    *oidc.Authenticator
}

const singleElementArrayLength = 1

// NewAPI creates and prepares a new API
func NewAPI(logger *zap.Logger, mainProcess *process.Process, tlsOptions orgtls.TLSOptions, configAPIAddress string, authenticator *oidc.Authenticator) (*API, error) {
	if mainProcess == nil {
		return nil, errors.New("process argument is nil. needed to close gracefully")
	}

	roots, orgKeyPair, err := orgtls.Load(tlsOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load tls certs")
	}

	orgCert := orgKeyPair.Leaf

	if len(orgCert.Subject.Organization) != singleElementArrayLength {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if len(orgCert.Subject.Organization) != singleElementArrayLength {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if configAPIAddress == "" {
		return nil, errors.New("config API address is not configured")
	}

	if authenticator == nil {
		return nil, errors.New("authenticator is not configured")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for api", zap.String("api-organization-name", organizationName))

	i := &API{
		logger:           logger.With(zap.String("api-organization-name", organizationName)),
		roots:            roots,
		orgCertKeyPair:   orgKeyPair,
		configAPIAddress: configAPIAddress,
		process:          mainProcess,
		mux:              runtime.NewServeMux(),
		authenticator:    authenticator,
	}

	return i, nil
}
