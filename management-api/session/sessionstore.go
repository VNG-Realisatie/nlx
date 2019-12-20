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

// Sessionstore provides the interface for a session store
type Sessionstore interface {
	Middleware(next http.Handler) http.Handler
	Routes() chi.Router
}

// SessionstoreImpl persists sessions and provides convinient helpers
type SessionstoreImpl struct {
	accountRepo repositories.Account
	cookiestore *sessions.CookieStore
	logger      *zap.Logger
	options     SessionstoreOptions
}

// SessionstoreOptions provides the flags to configure the SessionstoreImpl
type SessionstoreOptions struct {
	SecretKey           string `long:"secret-key" env:"SECRET_KEY" description:"Secret key that is used for signing sessions" required:"true"`
	SessionCookieSecure bool   `long:"session-cookie-secure" env:"SESSION_COOKIE_SECURE" description:"Use 'secure' cookies"`
	SessionCookieMaxAge int    `long:"session-cookie-maxage" env:"SESSION_COOKIE_MAXAGE" default:"3600" description:"The lifetime of a session, in seconds"`
}

// NewSessionstoreImpl creates a new Sessionstore
func NewSessionstoreImpl(logger *zap.Logger, options SessionstoreOptions, accountRepo repositories.Account) *SessionstoreImpl {
	cookiestore := sessions.NewCookieStore([]byte(options.SecretKey))
	cookiestore.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		Secure:   options.SessionCookieSecure,
		MaxAge:   options.SessionCookieMaxAge,
	}

	return &SessionstoreImpl{
		logger:      logger,
		cookiestore: cookiestore,
		accountRepo: accountRepo,
		options:     options,
	}
}

// NewSession gathers all session info for a Request
func (s *SessionstoreImpl) NewSession(r *http.Request) *Impl {
	session, _ := s.cookiestore.Get(r, cookieName)

	return &Impl{
		sessionstore: s,
		session:      session,
		r:            r,
	}
}

func getSession(r *http.Request) Session {
	value := r.Context().Value(contextKey)
	if value != nil {
		return value.(Session)
	}

	return nil
}

// Middleware propagates a new Context with a Session
func (s *SessionstoreImpl) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := s.NewSession(r)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKey, session)))
	})
}

// Routes provides the routes for the Router
func (s *SessionstoreImpl) Routes() chi.Router {
	r := chi.NewRouter()

	csrfMiddleware := csrf.Protect([]byte(s.options.SecretKey), csrf.Secure(s.options.SessionCookieSecure), csrf.MaxAge(s.options.SessionCookieMaxAge))
	r.Use(csrfMiddleware)

	r.Get("/login", s.preLogin())
	r.Post("/login", s.login())
	r.Get("/logout", s.logout())

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

func (s SessionstoreImpl) login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &loginRequest{}
		s.logger.Info("login", zap.Any("request", request))
		if err := render.Bind(r, request); err != nil {
			s.logger.Warn("Failed to login", zap.Error(err))
			http.Error(w, http.StatusText(400), 400)
			return
		}

		if err := getSession(r).Login(w, request.Account.ID); err != nil {
			s.logger.Warn("Failed to login", zap.Error(err))
			http.Error(w, http.StatusText(400), 400)
			return
		}

		account, err := getSession(r).Account()
		if err != nil {
			s.logger.Warn("Failed to get account", zap.Error(err))
			http.Error(w, http.StatusText(500), 500)
			return
		}

		render.JSON(w, r, loginReponse{Username: account.Name, Role: account.Role})
	}
}

func (s SessionstoreImpl) logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := getSession(r).Logout(w)
		if err != nil {
			s.logger.Debug("Failed to logout", zap.Error(err))
			http.Error(w, http.StatusText(500), 500)
		}
	}
}

func (s *SessionstoreImpl) preLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-CSRF-Token", csrf.Token(r))
	}
}
