// nolint:dupl
package configservice

import (
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

// ConfigService handles all requests for the config api
type ConfigService struct {
	logger                      *zap.Logger
	configDatabase              ConfigDatabase
	mainProcess                 *process.Process
	directoryRegistrationClient registrationapi.DirectoryRegistrationClient
}

// New creates new ConfigService
func New(logger *zap.Logger, mainProcess *process.Process, directoryRegistrationClient registrationapi.DirectoryRegistrationClient, configDatabase ConfigDatabase) *ConfigService {
	return &ConfigService{
		configDatabase:              configDatabase,
		logger:                      logger,
		mainProcess:                 mainProcess,
		directoryRegistrationClient: directoryRegistrationClient,
	}
}
