package oidc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"

	"go.nlx.io/nlx/management-api/pkg/oidc/mock"
)

// client is a special client that does not follow redirects automatically
var client = &http.Client{
	CheckRedirect: func(r *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func TestOnlyAuthenticated(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock.NewMockOAuth2Config(ctrl)
	mockStore := mock.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.Logger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is only visible when authenticated.")
	})

	srv := httptest.NewServer(authenticator.OnlyAuthenticated(mockHandler))
	defer srv.Close()

	tests := map[string]struct {
		session        func() *sessions.Session
		expectedStatus int
		expectedBody   string
	}{
		"unauthorized_request_should_fail": {
			func() *sessions.Session {
				return &sessions.Session{}
			},
			http.StatusUnauthorized,
			"unauthorized request\n",
		},
		"should_return_unauthorized_when_expired": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:   "arbitrary-user-id",
					ExpiresAt: time.Now().Add(-10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusUnauthorized,
			"unauthorized request\n",
		},
		"should_return_authorized_when_not_expired": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:   "arbitrary-user-id",
					ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
		"should_return_unauthorized_when_used_before_not_before": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:   "arbitrary-user-id",
					NotBefore: time.Now().Add(10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusUnauthorized,
			"unauthorized request\n",
		},
		"should_return_authorized_when_used_after_not_before": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:   "arbitrary-user-id",
					NotBefore: time.Now().Add(-10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
		"should_return_unauthorized_when_used_before_issued": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:  "arbitrary-user-id",
					IssuedAt: time.Now().Add(10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusUnauthorized,
			"unauthorized request\n",
		},
		"should_return_authorized_when_used_after_issued": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:  "arbitrary-user-id",
					IssuedAt: time.Now().Add(-10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
		"should_return_authorized_for_successful_request_without_optional_claims": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject: "arbitrary-user-id",
				}

				return session
			},
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
		"should_return_authorized_for_successful_request_with_optional_claims": {
			func() *sessions.Session {
				session := sessions.NewSession(mockStore, "nlx_management_session")
				session.Values["claims"] = &Claims{
					Subject:   "arbitrary-user-id",
					IssuedAt:  time.Now().Add(-10 * time.Second).Unix(),
					NotBefore: time.Now().Add(-10 * time.Second).Unix(),
					ExpiresAt: time.Now().Add(10 * time.Second).Unix(),
				}

				return session
			},
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			mockStore.
				EXPECT().
				Get(gomock.Any(), "nlx_management_session").
				Return(tt.session(), nil)

			resp, err := client.Get(srv.URL)
			assert.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, string(bodyBytes))
		})
	}
}

func TestAuthenticateEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock.NewMockOAuth2Config(ctrl)
	mockStore := mock.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.Logger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	srv := httptest.NewServer(authenticator.Routes())
	defer srv.Close()

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(&sessions.Session{}, nil)
	mockOAuth2Config.EXPECT().AuthCodeURL("").Return("https://example.com/some-redirect-url")

	resp, err := client.Get(fmt.Sprintf("%s/authenticate", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Location"), "https://example.com/some-redirect-url")
	resp.Body.Close()
}

func TestMeEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock.NewMockOAuth2Config(ctrl)
	mockStore := mock.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.Logger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	srv := httptest.NewServer(authenticator.Routes())
	defer srv.Close()

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(&sessions.Session{}, nil)

	resp, err := client.Get(fmt.Sprintf("%s/me", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	resp.Body.Close()
}

func TestCallbackEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock.NewMockOAuth2Config(ctrl)
	mockStore := mock.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.Logger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	srv := httptest.NewServer(authenticator.Routes())
	defer srv.Close()

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(&sessions.Session{}, nil).AnyTimes()
	mockOAuth2Config.EXPECT().Exchange(gomock.Any(), "1337").Return(&oauth2.Token{}, nil)

	// this tests the non-happy trail
	resp, err := client.Get(fmt.Sprintf("%s/callback?code=1337", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Location"), "/")
	resp.Body.Close()
}

func TestLogoutEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock.NewMockOAuth2Config(ctrl)
	mockStore := mock.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.Logger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	srv := httptest.NewServer(authenticator.Routes())
	defer srv.Close()

	mockSession := sessions.NewSession(mockStore, "nlx_management_session")
	mockSession.Values["claims"] = &Claims{
		Subject: "42",
	}

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(mockSession, nil).AnyTimes()
	mockStore.EXPECT().Save(gomock.Any(), gomock.Any(), mockSession).Return(nil).AnyTimes()

	resp, err := client.Post(fmt.Sprintf("%s/logout", srv.URL), "application/x-www-form-urlencoded", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Location"), "/")

	_, claimsExists := mockSession.Values["claims"]
	assert.False(t, claimsExists)

	resp.Body.Close()
}
