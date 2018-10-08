package config

import (
	"strings"

	"go.nlx.io/nlx/insight-api/irma"

	"github.com/ktr0731/toml"
	"go.uber.org/zap"
)

// InsightConfig is the top-level for the insight configuration file.
type InsightConfig struct {
	DataSubjects map[string]DataSubject
}

type DataSubject struct {
	Label          string           `toml:"label"`
	IrmaAttributes []irma.Attribute `toml:"irma-attributes"`
}

// LoadInsightConfig reads the service config from disk and returns.
func LoadInsightConfig(logger *zap.Logger, insightConfigLocation string) *InsightConfig {
	insightConfig := &InsightConfig{}
	tomlMetaData, err := toml.DecodeFile(insightConfigLocation, insightConfig)
	if err != nil {
		logger.Fatal("failed to load service config", zap.Error(err))
	}
	if len(tomlMetaData.Undecoded()) > 0 {
		logger.Fatal("unsupported values in toml", zap.String("key", strings.Join(tomlMetaData.Undecoded()[0], ">")))
	}
	return insightConfig
}
