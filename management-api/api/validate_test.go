package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/api"
)

func TestServiceValidate(t *testing.T) {
	testService := api.Service{
		Name: "my-service",
	}

	err := testService.Validate()
	assert.Equal(t, errors.New("invalid endpoint URL for service my-service"), err)

	testService = api.Service{
		Name:        "my-service",
		EndpointURL: "my-service.test",
	}

	err = testService.Validate()
	assert.Equal(t, nil, err)
}

func TestInwayValidate(t *testing.T) {
	testInway := api.Inway{}

	err := testInway.Validate()
	assert.Equal(t, errors.New("invalid inway name"), err)

	testInway = api.Inway{
		Name: "inway42.test",
	}

	err = testInway.Validate()
	assert.Equal(t, nil, err)
}
