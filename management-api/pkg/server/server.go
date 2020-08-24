package server

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"

	"go.nlx.io/nlx/management-api/pkg/database"
)

// ManagementService handles all requests for the config api
type ManagementService struct {
	logger                      *zap.Logger
	configDatabase              database.ConfigDatabase
	mainProcess                 *process.Process
	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

// NewManagementService creates new ManagementService
func NewManagementService(logger *zap.Logger, mainProcess *process.Process, directoryRegistrationClient registrationapi.DirectoryRegistrationClient, configDatabase database.ConfigDatabase) *ManagementService {
	return &ManagementService{
		configDatabase:              configDatabase,
		logger:                      logger,
		mainProcess:                 mainProcess,
		directoryRegistrationClient: directoryRegistrationClient,
	}
}
