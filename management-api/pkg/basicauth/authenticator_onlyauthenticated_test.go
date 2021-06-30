// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth_test

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_OnlyAuthenticated(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		setup              func(authenticatorMocks)
		headers            map[string]string
		expectedStatusCode int
	}{
		"without_auth_header": {
			headers:            map[string]string{},
			expectedStatusCode: http.StatusUnauthorized,
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
			expectedStatusCode: http.StatusUnauthorized,
		},
		"when_database_call_fails": {
			headers: map[string]string{
				"Authorization": basicAuthorizationHeader("baz", "qux"),
			},
			setup: func(mocks authenticatorMocks) {
				mocks.configDatabase.
					EXPECT().
					VerifyUserCredentials(gomock.Any(), "baz", "qux").
					Return(false, errors.New("arbitrary error"))
			},
			expectedStatusCode: http.StatusUnauthorized,
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
			},
			expectedStatusCode: http.StatusOK,
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

			handler := authenticator.OnlyAuthenticated(Handler(customHandler))
			router.Method(http.MethodGet, "/", handler)

			srv := httptest.NewServer(router)
			defer srv.Close()

			client := &http.Client{}

			req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
			require.NoError(t, err)

			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			resp, err := client.Do(req)
			require.NoError(t, err)

			defer resp.Body.Close()

			require.Equal(t, tt.expectedStatusCode, resp.StatusCode)
		})
	}
}

func basicAuthorizationHeader(username, password string) string {
	auth := username + ":" + password
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))

	return fmt.Sprintf("Basic %s", encoded)
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func customHandler(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(http.StatusOK)

	return nil
}
