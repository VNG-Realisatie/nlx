// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package server

import (
	"time"

	"go.uber.org/zap"

	"go.nlx.io/nlx/txlog-api/api"
	"go.nlx.io/nlx/txlog-api/domain/txlog/storage"
)

type TXLogService struct {
	api.UnimplementedTXLogServer

	logger  *zap.Logger
	storage storage.Repository
	clock   Clock
}

type Clock interface {
	Now() time.Time
}

func NewTXLogService(logger *zap.Logger, s storage.Repository, clock Clock) *TXLogService {
	return &TXLogService{
		logger:  logger,
		storage: s,
		clock:   clock,
	}
}
