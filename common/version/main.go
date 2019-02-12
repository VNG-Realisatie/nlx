package version

import (
	"go.uber.org/zap"
)

var (
	// BuildVersion contains the version of the build
	BuildVersion = "unknown"
	// BuildSourceHash contains the git commit hash of the build
	BuildSourceHash = "unknown"
)

// Log writes the version and sourcehash to the logger
func Log(logger *zap.Logger) {
	logger.Info("Version information", zap.String("version", BuildVersion), zap.String("source-hash", BuildSourceHash))
}
