package server

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"

	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
)

// ManagementService handles all requests for the config api
type ManagementService struct {
	logger          *zap.Logger
	configDatabase  database.ConfigDatabase
	mainProcess     *process.Process
	directoryClient directory.Client
}

// New creates new ManagementService
func NewManagementService(logger *zap.Logger, mainProcess *process.Process, directoryClient directory.Client, configDatabase database.ConfigDatabase) *ManagementService {
	return &ManagementService{
		configDatabase:  configDatabase,
		logger:          logger,
		mainProcess:     mainProcess,
		directoryClient: directoryClient,
	}
}
