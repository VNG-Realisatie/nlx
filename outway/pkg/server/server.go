// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/outway/api"
)

type OutwayService struct {
	api.UnimplementedOutwayServer

	logger  *zap.Logger
	orgCert *tls.CertificateBundle
}

func NewOutwayService(logger *zap.Logger, orgCert *tls.CertificateBundle) *OutwayService {
	return &OutwayService{
		logger:  logger,
		orgCert: orgCert,
	}
}
