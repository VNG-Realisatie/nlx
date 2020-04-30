// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()

	handler := http.NewServeMux()
	handler.Handle("/", mux)

	server := httptest.NewServer(handler)
	client, _ := NewClient(server.URL)

	return client, mux, server.Close
}

func TestClientInvalidURL(t *testing.T) {
	_, err := NewClient(" http://endpoint")

	assert.Error(t, err)
}

func TestClientRequestHeaders(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "nlx-management", r.UserAgent(), "user agent does not match")
		assert.Contains(t, r.Header["Accept"], "application/json")
	})

	_, err := client.ListServices()
	assert.Error(t, err)
}

var listServicesTests = []struct {
	responseCode     int
	responseBody     string
	expectedServices []*Service
	expectError      bool
}{
	{
		200,
		"",
		nil,
		true,
	},
	{
		200,
		"{}",
		nil,
		false,
	},
	{
		200,
		`{"services": [ {"service_name": "test", "organization_name": "Test Corp"} ]}`,
		[]*Service{
			{
				Name:                 "test",
				OrganizationName:     "Test Corp",
				APISpecificationType: "",
				Inways:               nil,
			},
		},
		false,
	},
	{
		500,
		"",
		nil,
		true,
	},
}

func TestClientListServices(t *testing.T) {
	for i, test := range listServicesTests {
		name := strconv.Itoa(i + 1)
		test := test

		t.Run(name, func(t *testing.T) {
			client, mux, teardown := setup()
			defer teardown()

			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.responseCode)
				fmt.Fprint(w, test.responseBody)
			})

			services, err := client.ListServices()

			assert.Equal(t, test.expectedServices, services)

			if test.expectError {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
