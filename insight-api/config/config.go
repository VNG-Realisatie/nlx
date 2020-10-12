// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package config

import (
	"fmt"
	"strings"

	"github.com/ktr0731/toml"
	"go.uber.org/zap"
)

// InsightConfig is the top-level for the insight configuration file.
type InsightConfig struct {
	DataSubjects map[string]DataSubject
}

type DataSubject struct {
	Label          string          `toml:"label"`
	IrmaAttributes []IrmaAttribute `toml:"irma-attributes"`
}

type IrmaAttribute string

// LoadInsightConfig reads the service config from disk and returns.
func LoadInsightConfig(logger *zap.Logger, insightConfigLocation string) (*InsightConfig, error) {
	insightConfig := &InsightConfig{}

	tomlMetaData, err := toml.DecodeFile(insightConfigLocation, insightConfig)
	if err != nil {
		return nil, err
	}

	if len(tomlMetaData.Undecoded()) > 0 {
		return nil, fmt.Errorf("unsupported values in toml. key: " + strings.Join(tomlMetaData.Undecoded()[0], ">"))
	}

	return insightConfig, nil
}
