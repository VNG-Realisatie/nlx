// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package oidc

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"golang.org/x/oauth2"

	mock_auditlog "go.nlx.io/nlx/management-api/pkg/auditlog/mock"
	mock_database "go.nlx.io/nlx/management-api/pkg/database/mock"
	mock_oidc "go.nlx.io/nlx/management-api/pkg/oidc/mock"
)

// client is a special client that does not follow redirects automatically
var client = &http.Client{
	CheckRedirect: func(r *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

const testToken = "token"

//nolint:funlen // this is a test
func TestOnlyAuthenticated(t *testing.T) {
	tests := map[string]struct {
		setupMocks     func(store *mock_oidc.MockStore)
		expectedStatus int
		expectedHeader http.Header
	}{
		"when_no_session": {
			setupMocks: func(store *mock_oidc.MockStore) {
				store.
					EXPECT().
					Get(gomock.Any(), "nlx_management_session").
					Return(nil, errors.New("arbitrary error"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedHeader: http.Header{},
		},
		"when_no_token": {
			setupMocks: func(store *mock_oidc.MockStore) {
				store.
					EXPECT().
					Get(gomock.Any(), "nlx_management_session").
					Return(sessions.NewSession(store, "nlx_management_session"), nil)
			},
			expectedStatus: http.StatusUnauthorized,
			expectedHeader: http.Header{},
		},
		"happy_flow": {
			setupMocks: func(store *mock_oidc.MockStore) {
				session := sessions.NewSession(store, "nlx_management_session")

				session.Values[tokenName] = testToken

				store.
					EXPECT().
					Get(gomock.Any(), "nlx_management_session").
					Return(session, nil)
			},
			expectedStatus: http.StatusOK,
			expectedHeader: http.Header{
				"Authorization": []string{"Bearer token"},
			},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockOAuth2Config := mock_oidc.NewMockOAuth2Config(ctrl)
			mockStore := mock_oidc.NewMockStore(ctrl)

			authenticator := Authenticator{
				logger:       zaptest.NewLogger(t),
				oauth2Config: mockOAuth2Config,
				store:        mockStore,
			}

			var req *http.Request
			mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				req = r
				fmt.Fprintln(w, "unauthorized request")
			})

			srv := httptest.NewServer(authenticator.OnlyAuthenticated(mockHandler))
			defer srv.Close()

			tt.setupMocks(mockStore)

			resp, err := client.Get(srv.URL)
			assert.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			for header, value := range tt.expectedHeader {
				v, ok := req.Header[header]
				assert.True(t, ok)

				assert.Equal(t, value, v)
			}
		})
	}
}

func TestAuthenticateEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock_oidc.NewMockOAuth2Config(ctrl)
	mockStore := mock_oidc.NewMockStore(ctrl)
	mockDB := mock_database.NewMockConfigDatabase(ctrl)
	mockVerifier := mock_oidc.NewMockVerifier(ctrl)

	session := sessions.NewSession(mockStore, "nlx_management_session")
	session.Values[tokenName] = testToken

	authenticator := Authenticator{
		logger:       zaptest.NewLogger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
		db:           mockDB,
		oidcVerifier: mockVerifier,
	}

	router := chi.NewRouter()
	authenticator.MountRoutes(router)

	srv := httptest.NewServer(router)
	defer srv.Close()

	mockStore.
		EXPECT().
		Get(gomock.Any(), "nlx_management_session").
		Return(session, nil)
	mockVerifier.
		EXPECT().
		Verify(gomock.Any(), testToken).
		Return(nil, errors.New("arbitrary error"))
	mockStore.
		EXPECT().
		Save(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil)
	mockOAuth2Config.
		EXPECT().
		AuthCodeURL("").
		Return("https://example.com/some-redirect-url")

	resp, err := client.Get(fmt.Sprintf("%s/oidc/authenticate", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, "https://example.com/some-redirect-url", resp.Header.Get("Location"))
	resp.Body.Close()
}

func TestMeEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock_oidc.NewMockOAuth2Config(ctrl)
	mockStore := mock_oidc.NewMockStore(ctrl)

	authenticator := Authenticator{
		logger:       zaptest.NewLogger(t),
		oauth2Config: mockOAuth2Config,
		store:        mockStore,
	}

	router := chi.NewRouter()
	authenticator.MountRoutes(router)

	srv := httptest.NewServer(router)
	defer srv.Close()

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(&sessions.Session{}, nil)

	resp, err := client.Get(fmt.Sprintf("%s/oidc/me", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	resp.Body.Close()
}

func TestCallbackEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock_oidc.NewMockOAuth2Config(ctrl)
	mockStore := mock_oidc.NewMockStore(ctrl)

	auditLogger := mock_auditlog.NewMockLogger(ctrl)
	auditLogger.EXPECT().LoginFail(gomock.Any(), "Go-http-client/1.1")

	authenticator := Authenticator{
		logger:       zaptest.NewLogger(t),
		oauth2Config: mockOAuth2Config,
		auditLogger:  auditLogger,
		store:        mockStore,
	}

	router := chi.NewRouter()
	authenticator.MountRoutes(router)

	srv := httptest.NewServer(router)
	defer srv.Close()

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(&sessions.Session{}, nil).AnyTimes()
	mockOAuth2Config.EXPECT().Exchange(gomock.Any(), "1337").Return(&oauth2.Token{}, nil)

	// this tests the non-happy trail
	resp, err := client.Get(fmt.Sprintf("%s/oidc/callback?code=1337", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, "/login#auth-fail", resp.Header.Get("Location"))
	resp.Body.Close()
}

func TestLogoutEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOAuth2Config := mock_oidc.NewMockOAuth2Config(ctrl)
	mockStore := mock_oidc.NewMockStore(ctrl)

	auditLogger := mock_auditlog.NewMockLogger(ctrl)

	verifier := mock_oidc.NewMockVerifier(ctrl)
	verifier.EXPECT().Verify(gomock.Any(), gomock.Any()).Return(&oidc.IDToken{
		Issuer:          "",
		Audience:        nil,
		Subject:         "",
		Expiry:          time.Time{},
		IssuedAt:        time.Time{},
		Nonce:           "",
		AccessTokenHash: "",
	}, nil)

	authenticator := Authenticator{
		logger:       zaptest.NewLogger(t),
		oauth2Config: mockOAuth2Config,
		auditLogger:  auditLogger,
		store:        mockStore,
		oidcVerifier: verifier,
		getClaims: func(idToken *oidc.IDToken) (*IDTokenClaims, error) {
			return &IDTokenClaims{
				RegisteredClaims: jwt.RegisteredClaims{},
				Email:            "admin@example.com",
			}, nil
		},
	}

	router := chi.NewRouter()
	authenticator.MountRoutes(router)

	srv := httptest.NewServer(router)
	defer srv.Close()

	mockSession := sessions.NewSession(mockStore, "nlx_management_session")
	mockSession.Values["authorization"] = testToken

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(mockSession, nil).AnyTimes()
	mockStore.EXPECT().Save(gomock.Any(), gomock.Any(), mockSession).Return(nil).AnyTimes()

	auditLogger.EXPECT().LogoutSuccess(gomock.Any(), "admin@example.com", "Go-http-client/1.1")

	resp, err := client.Post(fmt.Sprintf("%s/oidc/logout", srv.URL), "application/x-www-form-urlencoded", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Location"), "/")

	_, claimsExists := mockSession.Values["claims"]
	assert.False(t, claimsExists)

	resp.Body.Close()
}
