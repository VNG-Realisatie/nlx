package config

import (
	"strings"
	"go.nlx.io/nlx/insight-api/irma"

	"github.com/ktr0731/toml"
	"go.uber.org/zap"
)

// ServiceConfig is the top-level for the service configuration file.
type ServiceConfig struct {
	DataSubjects map[string]DataSubject
}

type DataSubject struct {
    Label string                    `toml:"label"`
    IrmaAttributes []irma.Attribute `toml:"irma-attributes"`
}

// LoadServiceConfig reads the service config from disk and returns.
func LoadServiceConfig(logger *zap.Logger, serviceConfigLocation string) *ServiceConfig {
	serviceConfig := &ServiceConfig{}
	tomlMetaData, err := toml.DecodeFile(serviceConfigLocation, serviceConfig)
	if err != nil {
		logger.Fatal("failed to load service config", zap.Error(err))
	}
	if len(tomlMetaData.Undecoded()) > 0 {
		logger.Fatal("unsupported values in toml", zap.String("key", strings.Join(tomlMetaData.Undecoded()[0], ">")))
	}
	return serviceConfig
}
