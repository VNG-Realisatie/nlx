package process

import (
	"runtime"

	"go.uber.org/zap"
)

var logger *zap.Logger

// Setup performs common process setup for nlx daemons
func Setup(l *zap.Logger) {
	logger = l
	logger.Debug("setting up process")

	// Currently we develop for linux, binaries might be compiled for other platforms but we should warn the user about possible unexpected behaviour
	if runtime.GOOS != "linux" {
		logger.Warn("detected non-linux OS, program might behave unexpected", zap.String("os", runtime.GOOS))
	}

	setupSignals()
}
