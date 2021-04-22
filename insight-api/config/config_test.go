// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestLoadConfig(t *testing.T) {
	nop := zap.NewNop()
	_, err := LoadInsightConfig(nop, "../../testing/insight-api/insight-config-invalid.toml")
	assert.NotNil(t, err)
	assert.Equal(t, "unsupported values in toml. key: datasubjects>kenteken>unknown-attribute", err.Error())

	_, err = LoadInsightConfig(nop, "../../testing/insight-api/non-existing-config-file.toml")
	assert.NotNil(t, err)

	config, err := LoadInsightConfig(nop, "../../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	expectedDataSubjects := make(map[string]DataSubject)
	expectedDataSubjects["burgerservicenummer"] = DataSubject{
		Label: "Burgerservicenummer",
		IrmaAttributes: []IrmaAttribute{
			"irma-demo.nijmegen.bsn.bsn",
		},
	}
	expectedDataSubjects["kenteken"] = DataSubject{
		Label: "Kenteken",
		IrmaAttributes: []IrmaAttribute{
			"irma-demo.rvrd.vrn.vrn",
		},
	}
	assert.Equal(t, expectedDataSubjects, config.DataSubjects)
}
