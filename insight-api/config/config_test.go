// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package config

import (
	"testing"

	"go.nlx.io/nlx/insight-api/irma"

	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestLoadConfig(t *testing.T) {
	nop := zap.NewNop()
	_, err := LoadInsightConfig(nop, "../../testing/insight-api/insight-config-invalid.toml")
	assert.NotNil(t, err)
	assert.Equal(t, "unsupported values in toml. key: datasubjects>kenteken>unkown-attibute", err.Error())

	_, err = LoadInsightConfig(nop, "../../testing/insight-api/non-existing-config-file.toml")
	assert.NotNil(t, err)

	config, err := LoadInsightConfig(nop, "../../testing/insight-api/insight-config.toml")
	assert.Nil(t, err)

	expectedDataSubjects := make(map[string]DataSubject)
	expectedDataSubjects["burgerservicenummer"] = DataSubject{
		Label: "Burgerservicenummer",
		IrmaAttributes: []irma.Attribute{
			"irma-demo.nijmegen.bsn.bsn",
		},
	}
	expectedDataSubjects["kenteken"] = DataSubject{
		Label: "Kenteken",
		IrmaAttributes: []irma.Attribute{
			"irma-demo.rdw.vrn.vrn",
		},
	}
	assert.Equal(t, expectedDataSubjects, config.DataSubjects)
}
