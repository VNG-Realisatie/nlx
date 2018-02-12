// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

import (
	"crypto/x509"
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Inway handles incomming requests and holds a list of registered ServiceEndpoints.
// The Inway is responsible for selecting the correct ServiceEndpoint for an incomming request.
type Inway struct {
	logger           *zap.Logger
	organizationName string

	serviceEndpointsLock sync.RWMutex
	serviceEndpoints     map[string]ServiceEndpoint
}

// NewInway creates and prepares a new Inway.
func NewInway(l *zap.Logger, roots *x509.CertPool, orgCert *x509.Certificate) (*Inway, error) {
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}
	organizationName := orgCert.Subject.Organization[0]
	i := &Inway{
		logger:           l.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,

		serviceEndpoints: make(map[string]ServiceEndpoint),
	}
	return i, nil
}

// AddServiceEndpoint adds an ServiceEndpoint to the inway's internal registry.
func (i *Inway) AddServiceEndpoint(s ServiceEndpoint) error {
	i.serviceEndpointsLock.Lock()
	defer i.serviceEndpointsLock.Unlock()
	if _, exists := i.serviceEndpoints[s.ServiceName()]; exists {
		return errors.New("service endpoint for a service with the same name has already been registered")
	}
	i.serviceEndpoints[s.ServiceName()] = s
	return nil
}
