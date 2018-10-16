package process

import (
	"context"
	"runtime"

	"go.uber.org/zap"
)

var logger *zap.Logger

// Setup performs common process setup for nlx daemons
func Setup(l *zap.Logger) context.Context {
	logger = l
	logger.Debug("setting up process")

	// Currently we develop for linux, binaries might be compiled for other platforms but we should warn the user about possible unexpected behaviour
	if runtime.GOOS != "linux" {
		logger.Warn("detected non-linux OS, program might behave unexpected", zap.String("os", runtime.GOOS))
	}

	// TODO: would be better to provide context from caller to enable additional level of control infuture witout need to refactor all usages
	return setupSignals(context.Background())
}
