package oidc

import (
	"context"
	"encoding/gob"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

const cookieName = "nlx_management_session"

// OAuth2Config provides an interface for oauth2.Config
type OAuth2Config interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

// Provider provides an interface for oidc.Provider
type Provider interface {
	Endpoint() oauth2.Endpoint
	Verifier(config *oidc.Config) *oidc.IDTokenVerifier
}

// Store provides an interface for sessions.Store
type Store interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
	New(r *http.Request, name string) (*sessions.Session, error)
	Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error
}

// Authenticator implements the OIDC authentication mechanism
type Authenticator struct {
	logger       *zap.Logger
	oauth2Config OAuth2Config
	oidcConfig   *oidc.Config
	oidcProvider Provider
	store        Store
}

// User contains all the details of a specific user
type User struct {
	ID         string `json:"id"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	PictureURL string `json:"pictureUrl"`
}

// NewAuthenticator creates a new OIDC authenticator
func NewAuthenticator(logger *zap.Logger, options *Options) *Authenticator {
	gob.Register(&User{})

	store := sessions.NewCookieStore([]byte(options.SecretKey))
	store.Options = &sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   options.SessionCookieSecure,
		Path:     "/",
	}

	ctx := context.Background()
	oidcProvider, err := oidc.NewProvider(ctx, options.DiscoveryURL)

	if err != nil {
		logger.Fatal("could not initialize OIDC provider", zap.Error(err))
	}

	oauth2Config := &oauth2.Config{
		ClientID:     options.ClientID,
		ClientSecret: options.ClientSecret,
		RedirectURL:  options.RedirectURL,
		Endpoint:     oidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: options.ClientID,
	}

	return &Authenticator{
		logger:       logger,
		store:        store,
		oauth2Config: oauth2Config,
		oidcConfig:   oidcConfig,
		oidcProvider: oidcProvider,
	}
}

// Routes returns the OIDC routes
func (a *Authenticator) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/authenticate", a.authenticate)
	r.Get("/callback", a.callback)
	r.Post("/logout", a.logout)
	r.Get("/me", a.me)

	return r
}

// OnlyAuthenticated is middleware that only handles authenticated requests
func (a *Authenticator) OnlyAuthenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.GetUser(r) != nil {
			h.ServeHTTP(w, r)
			return
		}

		http.Error(w, "unauthorized request", http.StatusUnauthorized)
	})
}

// GetUser retrieves the User from a session
func (a *Authenticator) GetUser(r *http.Request) *User {
	session, _ := a.store.Get(r, cookieName)

	user, ok := session.Values["user"].(*User)
	if !ok {
		return nil
	}

	return user
}

func (a *Authenticator) authenticate(w http.ResponseWriter, r *http.Request) {
	if a.GetUser(r) != nil { // do not try to login again when user is already logged in
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Redirect(w, r, a.oauth2Config.AuthCodeURL(""), http.StatusFound)
}

func (a *Authenticator) me(w http.ResponseWriter, r *http.Request) {
	if user := a.GetUser(r); user != nil {
		render.JSON(w, r, user)
		return
	}

	http.Error(w, "unauthorized request", http.StatusUnauthorized)
}

func (a *Authenticator) callback(w http.ResponseWriter, r *http.Request) {
	if a.GetUser(r) != nil { // do not try to login again when user is already logged in
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ctx := context.Background()
	session, _ := a.store.Get(r, cookieName)

	oauth2Token, err := a.oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		a.logger.Info("could not parse id token")
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	verifier := a.oidcProvider.Verifier(a.oidcConfig)
	idToken, err := verifier.Verify(ctx, rawIDToken)

	if err != nil {
		a.logger.Info("could not verify id token", zap.Error(err))
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	var claims struct {
		Sub     string `json:"sub"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}

	err = idToken.Claims(&claims)

	if err != nil {
		a.logger.Info("could not extract claims", zap.Error(err))
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	session.Values["user"] = &User{
		ID:         claims.Sub,
		FullName:   claims.Name,
		Email:      claims.Email,
		PictureURL: claims.Picture,
	}

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r, cookieName)
	delete(session.Values, "user")

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
