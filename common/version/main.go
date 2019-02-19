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

// AddVersionToLogger adds the version and sourcehash to the logger
func AddVersionToLogger(logger *zap.Logger) *zap.Logger {
	return logger.With(zap.String("version", BuildVersion), zap.String("source-hash", BuildSourceHash))
}
