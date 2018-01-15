// Copyright 2018 VNG Realisatie. All rights reserved.
// Use of this source code is governed by the EUPL
// license that can be found in the LICENSE.md file.

package outway

import (
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Outway handles requests from inside the organization
type Outway struct {
	organizationName string // the organization running this outway

	logger *zap.Logger

	servicesLock sync.RWMutex
	services     map[string]*Service // services mapped by <organizationName>.<serviceName>, PoC shortcut in the absence of directory
}

// NewOutway creates a new Outway and sets it up to handle requests.
func NewOutway(l *zap.Logger, organizationName string) *Outway {
	i := &Outway{
		logger:           l.With(zap.String("outway-organization-name", organizationName)),
		organizationName: organizationName,

		services: make(map[string]*Service),
	}
	return i
}

// AddService adds a service and its inway to the outway's internal registry.
func (i *Outway) AddService(s *Service) error {
	i.servicesLock.Lock()
	defer i.servicesLock.Unlock()
	if _, exists := i.services[s.fullName()]; exists {
		return errors.New("service with same name has already been registered")
	}
	i.services[s.fullName()] = s
	return nil
}
