// nolint:dupl
package configservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

func TestNewConfigService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	configDatabase := ETCDConfigDatabase{}
	con, err := grpc.Dial("mock.directory.registration.api")
	assert.NotNil(t, err)
	directoryRegistrationClient := registrationapi.NewDirectoryRegistrationClient(con)

	expected := &ConfigService{
		logger:                      logger,
		mainProcess:                 testProcess,
		configDatabase:              configDatabase,
		directoryRegistrationClient: directoryRegistrationClient,
	}

	actual := New(logger, testProcess, directoryRegistrationClient, configDatabase)

	assert.Equal(t, expected, actual)
}
