// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package oidc

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
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
	store        Store
	oidcProvider Provider
	auditLogger  auditlog.Logger
	logger       *zap.Logger
	oauth2Config OAuth2Config
	oidcConfig   *oidc.Config
	db           database.ConfigDatabase
}

func NewAuthenticator(db database.ConfigDatabase, auditLogger auditlog.Logger, logger *zap.Logger, options *Options) *Authenticator {
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
		db:           db,
		auditLogger:  auditLogger,
		logger:       logger,
		store:        store,
		oauth2Config: oauth2Config,
		oidcConfig:   oidcConfig,
		oidcProvider: oidcProvider,
	}
}

func (a *Authenticator) MountRoutes(r chi.Router) {
	routes := chi.NewRouter()
	routes.Get("/authenticate", a.authenticate)
	routes.Get("/callback", a.callback)
	routes.Post("/logout", a.logout)
	routes.Get("/me", a.
		OnlyAuthenticated(http.HandlerFunc(a.me)).
		ServeHTTP,
	)

	r.Mount("/oidc", routes)
}

func (a *Authenticator) OnlyAuthenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := a.parseClaims(r)
		if err != nil {
			a.logger.Warn("authorization failed", zap.Error(err))

			http.Error(w, "unauthorized request", http.StatusUnauthorized)

			return
		}

		r.Header.Add("username", claims.User().FullName)

		h.ServeHTTP(w, r)
	})
}

func (a *Authenticator) parseClaims(r *http.Request) (*Claims, error) {
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
	if _, err := a.parseClaims(r); err == nil {
		http.Redirect(w, r, "/", http.StatusFound)

		return
	}

	// Remove previous claims from the session
	session, _ := a.store.Get(r, cookieName)

	if _, ok := session.Values["claims"]; ok {
		delete(session.Values, "claims")

		err := session.Save(r, w)
		if err != nil {
			http.Error(w, "unable to save session", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, a.oauth2Config.AuthCodeURL(""), http.StatusFound)
}

func (a *Authenticator) me(w http.ResponseWriter, r *http.Request) {
	claims, err := a.parseClaims(r)
	if err != nil {
		http.Error(w, "unauthorized request", http.StatusUnauthorized)

		return
	}

	render.JSON(w, r, claims.User())
}

func (a *Authenticator) callback(w http.ResponseWriter, r *http.Request) {
	err := func() error {
		ctx := r.Context()
		session, _ := a.store.Get(r, cookieName)

		oauth2Token, err := a.oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			return err
		}

		token, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			return errors.New("could not parse id token")
		}

		verifier := a.oidcProvider.Verifier(a.oidcConfig)

		idToken, err := verifier.Verify(ctx, token)
		if err != nil {
			return errors.New("could not verify id token")
		}

		// first, load the claims
		claims := &Claims{}

		if err = idToken.Claims(claims); err != nil {
			return errors.New("could not extract claims from ID-token")
		}

		// then, check if the user exists in the NLX Management database
		user, err := a.db.GetUser(ctx, claims.Email)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		if !user.HasRole(database.AdminRole) {
			return fmt.Errorf("user doesn't have the admin role")
		}

		session.Values["claims"] = claims

		if err := session.Save(r, w); err != nil {
			return fmt.Errorf("failed to save session: %w", err)
		}

		err = a.auditLogger.LoginSuccess(ctx, claims.User().FullName, r.Header.Get("User-Agent"))
		if err != nil {
			a.logger.Error("error writing to audit log", zap.Error(err))
		}

		return nil
	}()
	if err != nil {
		a.logger.Info("authentication failed", zap.Error(err))

		err = a.auditLogger.LoginFail(r.Context(), r.Header.Get("User-Agent"))
		if err != nil {
			a.logger.Error("error writing to audit log", zap.Error(err))
		}

		if errors.Is(err, database.ErrNotFound) {
			http.Redirect(w, r, "/login#auth-missing-user", http.StatusFound)

			return
		}

		http.Redirect(w, r, "/login#auth-fail", http.StatusFound)

		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.store.Get(r, cookieName)
	claims := session.Values["claims"].(*Claims)

	err := a.auditLogger.LogoutSuccess(r.Context(), claims.User().FullName, r.Header.Get("User-Agent"))
	if err != nil {
		a.logger.Error("error writing to audit log", zap.Error(err))
	}

	delete(session.Values, "claims")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
