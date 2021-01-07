package oidc

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

var ErrUnauthenticated = errors.New("not authenticated")

const cookieName = "nlx_management_session"

type OAuth2Config interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

type Provider interface {
	Endpoint() oauth2.Endpoint
	Verifier(config *oidc.Config) *oidc.IDTokenVerifier
}

type Store interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
	New(r *http.Request, name string) (*sessions.Session, error)
	Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error
}

type Authenticator struct {
	logger       *zap.Logger
	oauth2Config OAuth2Config
	oidcConfig   *oidc.Config
	oidcProvider Provider
	store        Store
}

func NewAuthenticator(logger *zap.Logger, options *Options) *Authenticator {
	gob.Register(&Claims{})

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

func (a *Authenticator) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/authenticate", a.authenticate)
	r.Get("/callback", a.callback)
	r.Post("/logout", a.logout)
	r.Get("/me", a.
		OnlyAuthenticated(http.HandlerFunc(a.me)).
		ServeHTTP,
	)

	return r
}

func (a *Authenticator) OnlyAuthenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := a.ParseClaims(r); err != nil {
			http.Error(w, "unauthorized request", http.StatusUnauthorized)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func (a *Authenticator) ParseClaims(r *http.Request) (*Claims, error) {
	session, _ := a.store.Get(r, cookieName)

	claims, ok := session.Values["claims"].(*Claims)
	if !ok {
		return nil, ErrUnauthenticated
	}

	if err := claims.Verify(); err != nil {
		return nil, err
	}

	return claims, nil
}

func (a *Authenticator) authenticate(w http.ResponseWriter, r *http.Request) {
	// Don't login again if the current user is still valid
	if _, err := a.ParseClaims(r); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	http.Redirect(w, r, a.oauth2Config.AuthCodeURL(""), http.StatusFound)
}

func (a *Authenticator) me(w http.ResponseWriter, r *http.Request) {
	claims, err := a.ParseClaims(r)
	if err != nil {
		http.Error(w, "unauthorized request", http.StatusUnauthorized)

		return
	}

	render.JSON(w, r, claims.User())
}

func (a *Authenticator) callback(w http.ResponseWriter, r *http.Request) {
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

	claims := &Claims{}

	if err := idToken.Claims(claims); err != nil {
		a.logger.Info("could not extract claims", zap.Error(err))
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	session.Values["claims"] = claims

	if err := session.Save(r, w); err != nil {
		http.Error(w, "failed to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r, cookieName)
	delete(session.Values, "claims")

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
