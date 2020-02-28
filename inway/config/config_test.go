package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	serviceConfig string
}

var tests = []struct {
	name              string
	args              args
	want              *ServiceConfig
	wantDeprecatedErr bool
}{
	{
		name: "empty",
		args: args{
			serviceConfig: `
`,
		},
		wantDeprecatedErr: true,
		want:              &ServiceConfig{},
	},
	{
		name: "broken",
		args: args{
			serviceConfig: `
broken
`,
		},
		want: nil,
	},
	{
		name: "v1_without_version",
		args: args{
			serviceConfig: `
[services]
  [services.test]
	endpoint-url = "https://example.com/endpoint"
	authorization-model = "whitelist"
	authorization-whitelist = ["test org"]
	documentation-url = "https://example.com/documentation"
	api-specification-document-url = "https://example.com/api-specification-document"
	irma-api-url = "https://example.com/irma-api"
	insight-api-url = "https://example.com/insight-api"
	ca-cert-path = "ca-cert-path"
	public-support-contact = "public-support@example.com"
	tech-support-contact = "tech-support@example.com"
	internal = false
`,
		},
		wantDeprecatedErr: true,
		want: &ServiceConfig{
			Services: map[string]ServiceDetails{
				"test": {ServiceDetailsBase: ServiceDetailsBase{
					EndpointURL:                 "https://example.com/endpoint",
					AuthorizationModel:          "whitelist",
					DocumentationURL:            "https://example.com/documentation",
					APISpecificationDocumentURL: "https://example.com/api-specification-document",
					InsightAPIURL:               "https://example.com/insight-api",
					IrmaAPIURL:                  "https://example.com/irma-api",
					CACertPath:                  "ca-cert-path",
					PublicSupportContact:        "public-support@example.com",
					TechSupportContact:          "tech-support@example.com",
					Internal:                    false,
				},
					AuthorizationWhitelist: []AuthorizationWhitelistItem{{
						OrganizationName: "test org",
					}},
				},
			},
		},
	},
	{
		name: "v1_with_version",
		args: args{
			serviceConfig: `
version = 1
[services]
  [services.test]
	endpoint-url = "https://example.com/endpoint"
	authorization-model = "whitelist"
	authorization-whitelist = ["test org"]
	documentation-url = "https://example.com/documentation"
	api-specification-document-url = "https://example.com/api-specification-document"
	irma-api-url = "https://example.com/irma-api"
	insight-api-url = "https://example.com/insight-api"
	ca-cert-path = "ca-cert-path"
	public-support-contact = "public-support@example.com"
	tech-support-contact = "tech-support@example.com"
	internal = false
`,
		},
		wantDeprecatedErr: true,
		want: &ServiceConfig{
			Services: map[string]ServiceDetails{
				"test": {ServiceDetailsBase: ServiceDetailsBase{
					EndpointURL:                 "https://example.com/endpoint",
					AuthorizationModel:          "whitelist",
					DocumentationURL:            "https://example.com/documentation",
					APISpecificationDocumentURL: "https://example.com/api-specification-document",
					InsightAPIURL:               "https://example.com/insight-api",
					IrmaAPIURL:                  "https://example.com/irma-api",
					CACertPath:                  "ca-cert-path",
					PublicSupportContact:        "public-support@example.com",
					TechSupportContact:          "tech-support@example.com",
					Internal:                    false,
				},
					AuthorizationWhitelist: []AuthorizationWhitelistItem{{
						OrganizationName: "test org",
					}},
				},
			},
		},
	},
	{
		name: "v2",
		args: args{
			serviceConfig: `
version = 2
[services]
  [services.test]
	endpoint-url = "https://example.com/endpoint"
	authorization-model = "whitelist"
	documentation-url = "https://example.com/documentation"
	api-specification-document-url = "https://example.com/api-specification-document"
	irma-api-url = "https://example.com/irma-api"
	insight-api-url = "https://example.com/insight-api"
	ca-cert-path = "ca-cert-path"
	public-support-contact = "public-support@example.com"
	tech-support-contact = "tech-support@example.com"
	internal = false
	[[services.test.authorization-whitelist]]
	  organization-name = "test org"
      public-key-hash = "finger"
`,
		},
		wantDeprecatedErr: false,
		want: &ServiceConfig{
			serviceConfigVersion: serviceConfigVersion{Version: ServiceConfigVersionV2},
			Services: map[string]ServiceDetails{
				"test": {ServiceDetailsBase: ServiceDetailsBase{
					EndpointURL:                 "https://example.com/endpoint",
					AuthorizationModel:          "whitelist",
					DocumentationURL:            "https://example.com/documentation",
					APISpecificationDocumentURL: "https://example.com/api-specification-document",
					InsightAPIURL:               "https://example.com/insight-api",
					IrmaAPIURL:                  "https://example.com/irma-api",
					CACertPath:                  "ca-cert-path",
					PublicSupportContact:        "public-support@example.com",
					TechSupportContact:          "tech-support@example.com",
					Internal:                    false,
				},
					AuthorizationWhitelist: []AuthorizationWhitelistItem{{
						OrganizationName: "test org",
						PublicKeyHash:    "finger",
					}},
				},
			},
		},
	},
	{
		name: "v2_with_version_old_whitelist",
		args: args{
			serviceConfig: `
version = 2
[services]
  [services.test]
	endpoint-url = "https://example.com/endpoint"
	authorization-model = "whitelist"
	authorization-whitelist = ["test org"]
	documentation-url = "https://example.com/documentation"
	api-specification-document-url = "https://example.com/api-specification-document"
	irma-api-url = "https://example.com/irma-api"
	insight-api-url = "https://example.com/insight-api"
	ca-cert-path = "ca-cert-path"
	public-support-contact = "public-support@example.com"
	tech-support-contact = "tech-support@example.com"
	internal = false
`,
		},
		wantDeprecatedErr: true,
		want:              nil,
	},
	{
		name: "unknown_field",
		args: args{
			serviceConfig: `
version = 2
[services]
  [services.test]
	unknown_url = "https://example.com/unknown"
	endpoint-url = "https://example.com/endpoint"
	authorization-model = "whitelist"
	authorization-whitelist = ["test org"]
	documentation-url = "https://example.com/documentation"
	api-specification-document-url = "https://example.com/api-specification-document"
	irma-api-url = "https://example.com/irma-api"
	insight-api-url = "https://example.com/insight-api"
	ca-cert-path = "ca-cert-path"
	public-support-contact = "public-support@example.com"
	tech-support-contact = "tech-support@example.com"
	internal = false
`,
		},
		wantDeprecatedErr: true,
		want:              nil,
	},
}

func TestLoadServiceConfig(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			file, err := ioutil.TempFile("", fmt.Sprintf("service-config.toml.%s.", tt.name))
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(file.Name())

			_, err = file.WriteString(tt.args.serviceConfig)
			if err != nil {
				log.Fatal(err)
			}
			got, err := LoadServiceConfig(file.Name())
			if tt.want == nil {
				assert.Errorf(t, err, "got=%v", got)
			} else {
				if tt.wantDeprecatedErr {
					assert.Errorf(t, err, "got=%v", got)
				} else {
					assert.NoError(t, err)
				}
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
