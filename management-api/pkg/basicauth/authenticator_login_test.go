// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setup              func(authenticatorMocks)
		formData           url.Values
		expectedStatusCode int
	}{
		"with_empty_form_data": {
			expectedStatusCode: http.StatusBadRequest,
		},
		"without_email": {
			formData: url.Values{
				"password": {"my-secure-password"},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"without_password": {
			formData: url.Values{
				"email": {"hoi@nlx.io"},
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"with_invalid_credentials": {
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "hoi@nlx.io", "password").
					Return(false, errors.New("arbitrary error"))
			},
			formData: url.Values{
				"email":    {"hoi@nlx.io"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		"happy_flow": {
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "hoi@nlx.io", "password").
					Return(true, nil)
			},
			formData: url.Values{
				"email":    {"hoi@nlx.io"},
				"password": {"password"},
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	for name, tt := range testCases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			authenticator, mocks := newAuthenticator(t)

			if tt.setup != nil {
				tt.setup(mocks)
			}

			router := chi.NewRouter()
			authenticator.MountRoutes(router)

			srv := httptest.NewServer(router)
			defer srv.Close()

			resp, err := http.PostForm(fmt.Sprintf("%s%s", srv.URL, "/basic-auth/login"), tt.formData)
			require.NoError(t, err)

			defer resp.Body.Close()

			require.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}
