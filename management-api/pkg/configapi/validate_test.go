package configapi_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/configapi"
)

func TestServiceValidate(t *testing.T) {
	testService := configapi.Service{
		Name:                  "my-service",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "none"},
	}

	err := testService.Validate()
	assert.Equal(t, errors.New("invalid endpoint URL for service my-service"), err)

	testService = configapi.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	err = testService.Validate()
	assert.Equal(t, errors.New("invalid authorization settings for service my-service"), err)

	testService = configapi.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "nonexisting"},
	}

	err = testService.Validate()
	assert.Equal(t, errors.New("invalid authorization mode for service my-service, expected whitelist or none, got nonexisting"), err)

	testService = configapi.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "whitelist"},
	}

	err = testService.Validate()
	assert.Equal(t, nil, err)

	testService = configapi.Service{
		Name:                  "my-service",
		EndpointURL:           "my-service.test",
		AuthorizationSettings: &configapi.Service_AuthorizationSettings{Mode: "whitelist", Authorizations: []*configapi.Service_AuthorizationSettings_Authorization{}},
	}

	err = testService.Validate()
	assert.Equal(t, nil, err)
}

func TestInwayValidate(t *testing.T) {
	testInway := configapi.Inway{}

	err := testInway.Validate()
	assert.Equal(t, errors.New("invalid inway name"), err)

	testInway = configapi.Inway{
		Name: "inway42.test",
	}

	err = testInway.Validate()
	assert.Equal(t, nil, err)
}
