// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package session

import (
	"context"
	"errors"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/models"
	"go.nlx.io/nlx/management-api/repositories"
)

type key string

const contextKey key = "session"
const cookieName string = "nlx_management_session"

// AuthenticationManager persists sessions and provides convinient helpers
type AuthenticationManager struct {
	accountRepo repositories.Account
	cookiestore *sessions.CookieStore
	logger      *zap.Logger
	options     AuthenticationManagerOptions
}

// AuthenticationManagerOptions provides the flags to configure the AuthenticationManager
type AuthenticationManagerOptions struct {
	SecretKey           string `long:"secret-key" env:"SECRET_KEY" description:"Secret key that is used for signing sessions" required:"true"`
	SessionCookieSecure bool   `long:"session-cookie-secure" env:"SESSION_COOKIE_SECURE" description:"Use 'secure' cookies"`
	SessionCookieMaxAge int    `long:"session-cookie-maxage" env:"SESSION_COOKIE_MAXAGE" default:"3600" description:"The lifetime of a session, in seconds"`
}

// NewAuthenticationManager creates a new AuthenticationManager
func NewAuthenticationManager(logger *zap.Logger, options AuthenticationManagerOptions, accountRepo repositories.Account) *AuthenticationManager {
	cookiestore := sessions.NewCookieStore([]byte(options.SecretKey))
	cookiestore.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		Secure:   options.SessionCookieSecure,
		MaxAge:   options.SessionCookieMaxAge,
	}

	return &AuthenticationManager{
		logger:      logger,
		cookiestore: cookiestore,
		accountRepo: accountRepo,
		options:     options,
	}
}

// NewHttpSession gathers all session info for a Request
func (am *AuthenticationManager) NewHTTPSession(r *http.Request) *HTTPSession {
	session, _ := am.cookiestore.Get(r, cookieName)

	return &HTTPSession{
		authenicationManager: am,
		session:              session,
		r:                    r,
	}
}

func getSession(r *http.Request) Session {
	value := r.Context().Value(contextKey)
	if value != nil {
		session, ok := value.(Session)
		if ok {
			return session
		}
	}

	return nil
}

// Middleware propagates a new Context with a Session
func (am *AuthenticationManager) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := am.NewHTTPSession(r)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKey, session)))
	})
}

// Routes provides the routes for the Router
func (am *AuthenticationManager) Routes() chi.Router {
	r := chi.NewRouter()

	csrfMiddleware := csrf.Protect([]byte(am.options.SecretKey), csrf.Secure(am.options.SessionCookieSecure), csrf.MaxAge(am.options.SessionCookieMaxAge))
	r.Use(csrfMiddleware)

	r.Get("/login", am.preLogin())
	r.Post("/login", am.login())
	r.Get("/logout", am.logout())

	return r
}

type loginRequest struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Account  *models.Account `json:"-"`
}

// Bind validates the loginRequest and sets the Account when the username and password match the one stored in the repository
func (request *loginRequest) Bind(r *http.Request) error {
	if request.Username == "" {
		return errors.New("username is mandatory")
	}

	if request.Password == "" {
		return errors.New("password is mandatory")
	}

	account, err := getSession(r).AccountByName(request.Username)
	if err != nil {
		return err
	}

	match, err := argon2id.ComparePasswordAndHash(request.Password, account.PasswordHash)
	if err != nil {
		return err
	}

	if !match {
		return errors.New("password does not match stored hash")
	}

	request.Account = account

	return nil
}

type loginReponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (am AuthenticationManager) login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &loginRequest{}
		am.logger.Info("login", zap.Any("request", request))

		if err := render.Bind(r, request); err != nil {
			am.logger.Warn("Failed to login", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		if err := getSession(r).Login(w, request.Account.ID); err != nil {
			am.logger.Warn("Failed to login", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		account, err := getSession(r).Account()
		if err != nil {
			am.logger.Warn("Failed to get account", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			return
		}

		render.JSON(w, r, loginReponse{Username: account.Name, Role: account.Role})
	}
}

func (am AuthenticationManager) logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := getSession(r).Logout(w)
		if err != nil {
			am.logger.Debug("Failed to logout", zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (am *AuthenticationManager) preLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-CSRF-Token", csrf.Token(r))
		session := getSession(r)
		authenticated, err := session.IsAuthenticated()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		if authenticated {
			account, err := session.Account()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			render.JSON(w, r, loginReponse{Username: account.Name, Role: account.Role})
		}
	}
}
