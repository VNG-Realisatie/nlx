// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/management-api/pkg/database"
)

//nolint:funlen // this is a test method
func Test_Me(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setup          func(authenticatorMocks)
		headers        map[string]string
		expectedStatus int
		expectedBody   string
	}{
		"without_auth_header": {
			headers:        map[string]string{},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		"with_invalid_credentials": {
			headers: map[string]string{
				"Authorization": basicAuthorizationHeader("foo", "bar"),
			},
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "foo", "bar").
					Return(false, nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Unauthorized\n",
		},
		"when_verify_credentials_database_call_fails": {
			headers: map[string]string{
				"Authorization": basicAuthorizationHeader("foo", "bar"),
			},
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "foo", "bar").
					Return(false, errors.New("arbitrary error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error\n",
		},
		"when_get_user_database_call_fails": {
			headers: map[string]string{
				"Authorization": basicAuthorizationHeader("foo", "bar"),
			},
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "foo", "bar").
					Return(true, nil)

				mocks.configDatabase.
					EXPECT().
					GetUser(gomock.Any(), "foo").
					Return(nil, errors.New("arbitrary error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal Server Error\n",
		},
		"happy_flow": {
			headers: map[string]string{
				"Authorization": basicAuthorizationHeader("foo", "bar"),
			},
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "foo", "bar").
					Return(true, nil)

				mocks.configDatabase.
					EXPECT().
					GetUser(gomock.Any(), "foo").
					Return(&database.User{
						ID:       42,
						Email:    "foo@bar.com",
						Password: "my-password",
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "{\"id\":\"42\",\"fullName\":\"\",\"email\":\"foo@bar.com\",\"pictureUrl\":\"\"}\n",
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

			client := &http.Client{}

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/basic-auth/me", srv.URL), nil)
			require.NoError(t, err)

			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			resp, err := client.Do(req)
			require.NoError(t, err)

			defer resp.Body.Close()

			require.Equal(t, tt.expectedStatus, resp.StatusCode)

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, string(bodyBytes))
		})
	}
}
