package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplateForSubdomain(t *testing.T) {
	tests := map[string]struct {
		Subdomain        string
		ExpectedTemplate string
	}{
		"review": {
			Subdomain:        "review",
			ExpectedTemplate: "templates/sites-review.html",
		},
		"acc": {
			Subdomain:        "acc",
			ExpectedTemplate: "templates/sites-acc.html",
		},
		"pre-prod": {
			Subdomain:        "pre-prod",
			ExpectedTemplate: "templates/sites.html",
		},
		"prod": {
			Subdomain:        "prod",
			ExpectedTemplate: "templates/sites.html",
		},
	}

	for name, test := range tests {
		tc := test

		t.Run(name, func(t *testing.T) {
			template := templateForSubdomain(tc.Subdomain)
			assert.Equal(t, tc.ExpectedTemplate, template)
		})
	}
}
