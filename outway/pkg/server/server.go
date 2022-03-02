// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"crypto"

	"go.uber.org/zap"

	"go.nlx.io/nlx/common/delegation"

	"go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/outway/api"
)

type SignFunction func(privateKey crypto.PrivateKey, claims *delegation.JWTClaims) (string, error)

type OutwayService struct {
	api.UnimplementedOutwayServer

	logger  *zap.Logger
	orgCert *tls.CertificateBundle

	signFunction SignFunction
}

func NewOutwayService(logger *zap.Logger, orgCert *tls.CertificateBundle, signFunction SignFunction) *OutwayService {
	return &OutwayService{
		logger:       logger,
		orgCert:      orgCert,
		signFunction: signFunction,
	}
}
