// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/domain"
)

func Test_NewTermsOfServiceStatus(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args    *domain.NewTermsOfServiceStatusArgs
		wantErr string
	}{
		"invalid_username": {
			args: &domain.NewTermsOfServiceStatusArgs{
				Username:  "",
				CreatedAt: time.Now(),
			},
			wantErr: "Username: cannot be blank.",
		},
		"invalid_created_at": {
			args: &domain.NewTermsOfServiceStatusArgs{
				Username:  "admin",
				CreatedAt: time.Now().Add(time.Hour),
			},
			wantErr: "CreatedAt: must not be in the future.",
		},
		"happy_flow": {
			args: &domain.NewTermsOfServiceStatusArgs{
				Username:  "admin",
				CreatedAt: time.Now(),
			},
			wantErr: "",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			got, err := domain.NewTermsOfServiceStatus(tt.args)

			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.Nil(t, err)

				require.Equal(t, tt.args.Username, got.Username())
				require.Equal(t, tt.args.CreatedAt, got.CreatedAt())
			}
		})
	}
}
