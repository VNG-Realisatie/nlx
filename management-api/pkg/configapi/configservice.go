package configapi

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"

	"go.nlx.io/nlx/management-api/pkg/database"
)

// ConfigService handles all requests for the config api
type ConfigService struct {
	logger                      *zap.Logger
	configDatabase              database.ConfigDatabase
	mainProcess                 *process.Process
	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

// New creates new ConfigService
func New(logger *zap.Logger, mainProcess *process.Process, directoryRegistrationClient registrationapi.DirectoryRegistrationClient, configDatabase database.ConfigDatabase) *ConfigService {
	return &ConfigService{
		configDatabase:              configDatabase,
		logger:                      logger,
		mainProcess:                 mainProcess,
		directoryRegistrationClient: directoryRegistrationClient,
	}
}
