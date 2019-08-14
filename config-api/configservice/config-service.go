// nolint:dupl
package configservice

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
)

// ConfigService handles all requests for the config api
type ConfigService struct {
	logger         *zap.Logger
	configDatabase ConfigDatabase
	mainProcess    *process.Process
}

// New creates new ConfigService
func New(logger *zap.Logger, mainProcess *process.Process, configDatabase ConfigDatabase) *ConfigService {
	return &ConfigService{
		configDatabase: configDatabase,
		logger:         logger,
		mainProcess:    mainProcess,
	}
}
