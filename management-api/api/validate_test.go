package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/api"
)

func TestServiceValidate(t *testing.T) {
	testService := api.Service{
		Name: "my-service",
	}

	err := testService.Validate()
	assert.Equal(
		t,
		"endpointURL: cannot be blank.",
		err.Error(),
	)

	testService = api.Service{
		Name:        "my-service",
		EndpointURL: "https://my-service.test",
	}

	err = testService.Validate()
	assert.Equal(t, nil, err)
}

func TestInwayValidate(t *testing.T) {
	testInway := api.Inway{}

	err := testInway.Validate()
	assert.Equal(t, "name: cannot be blank.", err.Error())

	testInway = api.Inway{
		Name: "inway42.test",
	}

	err = testInway.Validate()
	assert.Equal(t, nil, err)
}
