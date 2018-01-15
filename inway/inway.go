// Copyright 2018 VNG Realisatie. All rights reserved.
// Use of this source code is governed by the EUPL
// license that can be found in the LICENSE.md file.

package inway

import (
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
func NewInway(l *zap.Logger, organizationName string) *Inway {
	i := &Inway{
		logger:           l.With(zap.String("inway-organization-name", organizationName)),
		organizationName: organizationName,

		serviceEndpoints: make(map[string]ServiceEndpoint),
	}
	return i
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
