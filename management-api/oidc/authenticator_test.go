package oidc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fgrosse/zaptest"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"

	"go.nlx.io/nlx/management-api/oidc/mock"
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

	mockLoggedInSession := sessions.NewSession(mockStore, "nlx_management_session")
	mockLoggedInSession.Values["user"] = &User{
		Sub: "42",
	}

	srv := httptest.NewServer(authenticator.OnlyAuthenticated(mockHandler))
	defer srv.Close()

	tests := []struct {
		session        *sessions.Session
		expectedStatus int
		expectedBody   string
	}{
		{
			&sessions.Session{},
			http.StatusUnauthorized,
			"unauthorized request\n",
		},
		{
			mockLoggedInSession,
			http.StatusOK,
			"This is only visible when authenticated.\n",
		},
	}

	for _, test := range tests {
		mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(test.session, nil)

		resp, err := client.Get(srv.URL)
		assert.NoError(t, err)

		defer resp.Body.Close()

		assert.Equal(t, test.expectedStatus, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedBody, string(bodyBytes))
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
	mockSession.Values["user"] = &User{
		Sub: "42",
	}

	mockStore.EXPECT().Get(gomock.Any(), "nlx_management_session").Return(mockSession, nil).AnyTimes()
	mockStore.EXPECT().Save(gomock.Any(), gomock.Any(), mockSession).Return(nil).AnyTimes()

	resp, err := client.Get(fmt.Sprintf("%s/logout", srv.URL))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, resp.Header.Get("Location"), "/")

	_, userExists := mockSession.Values["user"]
	assert.False(t, userExists)

	resp.Body.Close()
}
