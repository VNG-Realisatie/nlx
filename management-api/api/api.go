// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package api

import (
	"crypto/tls"
	"crypto/x509"

	"go.nlx.io/nlx/management-api/authorization"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/management-api/session"
)

// API handles incoming requests, authenticates them and forwards them to the config-api and txlog-api
type API struct {
	logger           *zap.Logger
	roots            *x509.CertPool
	orgCertKeyPair   *tls.Certificate
	process          *process.Process
	mux              *runtime.ServeMux
	configAPIAddress string
	sessionstore     session.Sessionstore
	authorizer       authorization.Authorizer
}

// NewAPI creates and prepares a new API
func NewAPI(logger *zap.Logger, mainProcess *process.Process, tlsOptions orgtls.TLSOptions, configAPIAddress string, sessionstore session.Sessionstore, authorizer authorization.Authorizer) (*API, error) {
	if mainProcess == nil {
		return nil, errors.New("process argument is nil. needed to close gracefully")
	}

	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load tls certs")
	}

	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	orgCertKeyPair, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load x509 keypair for organization")
	}

	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}

	if configAPIAddress == "" {
		return nil, errors.New("config API address is not configured")
	}

	if sessionstore == nil {
		return nil, errors.New("sessionstore is not configured")
	}

	if authorizer == nil {
		return nil, errors.New("authorizer is not configured")
	}

	organizationName := orgCert.Subject.Organization[0]
	logger.Info("loaded certificates for api", zap.String("api-organization-name", organizationName))

	i := &API{
		logger:           logger.With(zap.String("api-organization-name", organizationName)),
		roots:            roots,
		orgCertKeyPair:   &orgCertKeyPair,
		configAPIAddress: configAPIAddress,
		process:          mainProcess,
		mux:              runtime.NewServeMux(),
		sessionstore:     sessionstore,
		authorizer:       authorizer,
	}

	return i, nil
}
