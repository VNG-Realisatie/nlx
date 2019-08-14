// nolint:dupl
package configservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/common/process"
	"go.uber.org/zap"
)

func TestNewConfigService(t *testing.T) {
	logger := zap.NewNop()
	testProcess := process.NewProcess(logger)
	configDatabase := ETCDConfigDatabase{}

	expected := &ConfigService{
		logger:         logger,
		mainProcess:    testProcess,
		configDatabase: configDatabase,
	}

	actual := New(logger, testProcess, configDatabase)

	assert.Equal(t, expected, actual)
}
