// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package certportal

import (
	"go.uber.org/zap"
)

type CertPortal struct {
	logger *zap.Logger
	caHost string
}

// NewCertPortal creates a new CertPortal and sets it up to handle requests.
func NewCertPortal(l *zap.Logger, caHost string) *CertPortal {
	i := &CertPortal{
		logger: l,
		caHost: caHost,
	}
	return i
}
