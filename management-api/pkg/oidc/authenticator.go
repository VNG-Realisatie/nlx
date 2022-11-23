// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package oidc

import (
	"context"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/oidc/pgsessionstore"
)

var ErrUnauthenticated = errors.New("not authenticated")

const (
	bearer          string = "bearer"
	cookieName      string = "nlx_management_session"
	tokenName       string = "authorization"
	cleanupInterval        = time.Minute * 5
	maxSessionSize         = 1024 * 1024 * 1 // 1 MB
)

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

type Verifier interface {
	Verify(ctx context.Context, token string) (*oidc.IDToken, error)
}

type GetClaimsFromTokenFunc func(idToken *oidc.IDToken) (*IDTokenClaims, error)

type Authenticator struct {
	store        Store
	oidcProvider Provider
	auditLogger  auditlog.Logger
	logger       *zap.Logger
	oauth2Config OAuth2Config
	oidcConfig   *oidc.Config
	oidcVerifier Verifier
	getClaims    GetClaimsFromTokenFunc
	db           database.ConfigDatabase
	close        func()
}

type IDTokenClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type OurUserInfoGetter struct{}

func NewAuthenticator(httpsessionsDB *sql.DB, db database.ConfigDatabase, auditLogger auditlog.Logger, logger *zap.Logger, provider *oidc.Provider, verifier Verifier, getClaims GetClaimsFromTokenFunc, options *Options) (*Authenticator, error) {
	gob.Register(&Claims{})

	store, err := pgsessionstore.New(logger, httpsessionsDB, []byte(options.SecretKey))
	if err != nil {
		return nil, err
	}

	store.MaxLength(maxSessionSize)

	// Run a background goroutine to clean up expired sessions from the database.
	quit, done := store.StartCleanup(cleanupInterval)

	closeFn := func() {
		store.StopCleanup(quit, done)
		store.Close()
	}

	store.Options = &sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   options.SessionCookieSecure,
		Path:     "/",
	}

	oauth2Config := &oauth2.Config{
		ClientID:     options.ClientID,
		ClientSecret: options.ClientSecret,
		RedirectURL:  options.RedirectURL,
		Endpoint:     provider.Endpoint(),

		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
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
		oidcProvider: provider,
		oidcVerifier: verifier,
		getClaims:    getClaims,
		close:        closeFn,
	}, nil
}

func (a *Authenticator) Close() {
	a.close()
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
		logger := a.logger.With(zap.String("handler", "OnlyAuthenticated"))

		session, err := a.store.Get(r, cookieName)
		if err != nil {
			logger.Debug("failed to get session from store", zap.Error(err))
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		token, ok := session.Values[tokenName].(string)
		if !ok {
			logger.Debug("failed to get token from session", zap.String("token-name", tokenName), zap.Any("session-values", session.Values))
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		r.Header.Add("Authorization", "Bearer "+token)

		h.ServeHTTP(w, r)
	})
}

func (a *Authenticator) authenticate(w http.ResponseWriter, r *http.Request) {
	logger := a.logger.With(zap.String("handler", "authenticate"))

	session, err := a.store.Get(r, cookieName)
	if err != nil {
		logger.Error("unable to get nlx management cookie from the store", zap.Error(err))
		http.Error(w, "unauthorized request", http.StatusUnauthorized)

		return
	}

	bearerToken, ok := session.Values[tokenName].(string)
	if ok {
		_, err := a.oidcVerifier.Verify(r.Context(), bearerToken)
		if err == nil {
			logger.Error("cannot verify token", zap.Error(err))

			http.Redirect(w, r, "/", http.StatusFound)

			return
		}
	}

	if _, ok := session.Values[tokenName]; ok {
		delete(session.Values, tokenName)

		err := session.Save(r, w)
		if err != nil {
			logger.Error("unable to save session", zap.Error(err))

			http.Error(w, "unable to save session", http.StatusInternalServerError)

			return
		}
	}

	http.Redirect(w, r, a.oauth2Config.AuthCodeURL(""), http.StatusFound)
}

func (a *Authenticator) me(w http.ResponseWriter, r *http.Request) {
	logger := a.logger.With(zap.String("handler", "me"))

	session, err := a.store.Get(r, cookieName)
	if err != nil {
		logger.Error("unable to get nlx management cookie from the store", zap.Error(err))

		http.Error(w, "unauthorized request", http.StatusUnauthorized)

		return
	}

	bearerToken, ok := session.Values[tokenName].(string)
	if !ok {
		logger.Debug("failed to get token from session", zap.String("token-name", tokenName))

		http.Error(w, "unauthorized request, could not get bearer token", http.StatusUnauthorized)
		return
	}

	idToken, err := a.oidcVerifier.Verify(r.Context(), bearerToken)
	if err != nil {
		logger.Error("cannot verify token", zap.Error(err))

		http.Error(w, "unauthorized request, invalid token", http.StatusUnauthorized)

		return
	}

	// first, load the claims
	claims := &Claims{}

	if err = idToken.Claims(claims); err != nil {
		logger.Error("unable to parse id token claims", zap.Error(err))

		http.Error(w, "unauthorized request, could not parse token", http.StatusUnauthorized)

		return
	}

	render.JSON(w, r, claims.User())
}

func (a *Authenticator) callback(w http.ResponseWriter, r *http.Request) {
	err := func() error {
		ctx := r.Context()

		session, err := a.store.Get(r, cookieName)
		if err != nil {
			return errors.New("could not get session from cookie")
		}

		accessToken, err := a.oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			return err
		}

		token, ok := accessToken.Extra("id_token").(string)
		if !ok {
			return errors.New("could not parse id token")
		}

		idToken, err := a.oidcVerifier.Verify(ctx, token)
		if err != nil {
			return errors.New("could not verify id token")
		}

		// first, load the claims
		claims := &Claims{}

		if err = idToken.Claims(claims); err != nil {
			return errors.New("could not extract claims from ID-token")
		}

		session.Values[tokenName] = accessToken.AccessToken

		// Convert Unix timestamp in claim to amount of seconds from now.
		// MaxAge needs seconds from now
		session.Options.MaxAge = int(time.Until(time.Unix(claims.ExpiresAt, 0)).Seconds())

		if err = session.Save(r, w); err != nil {
			return fmt.Errorf("failed to save session: %w", err)
		}

		err = a.auditLogger.LoginSuccess(ctx, claims.User().Email, r.Header.Get("User-Agent"))
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
	session, err := a.store.Get(r, cookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	bearerToken := session.Values[tokenName].(string)

	token, err := a.oidcVerifier.Verify(r.Context(), bearerToken)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	claim, err := a.getClaims(token)
	if err != nil {
		a.logger.Error("unable to parse claims", zap.Error(err))
		http.Error(w, "unable to parse claims", http.StatusInternalServerError)

		return
	}

	err = a.auditLogger.LogoutSuccess(r.Context(), claim.Email, r.Header.Get("User-Agent"))
	if err != nil {
		a.logger.Error("error writing to audit log", zap.Error(err))
	}

	delete(session.Values, tokenName)

	err = session.Save(r, w)
	if err != nil {
		a.logger.Error("unable to save session", zap.Error(err))

		http.Error(w, "unable to save session", http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *Authenticator) UnaryServerInterceptor(configDatabase database.ConfigDatabase, getUserFromDatabase func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not extract metadata from context")
		}

		authHeader := md.Get("Authorization")

		// Outway and Inway do not send the authorization header
		if len(authHeader) == 0 {
			return handler(ctx, req)
		}

		var userAgent string
		if len(md.Get("grpcgateway-user-agent")) > 0 {
			userAgent = md.Get("grpcgateway-user-agent")[0]
		} else if len(md.Get("user-agent")) > 0 {
			userAgent = md.Get("user-agent")[0]
		}

		bearerToken, ok := extractTokenFromAuthHeader(authHeader[0])
		if !ok {
			return "", status.Errorf(codes.Unauthenticated, "could not extract bearer token from Authorization header")
		}

		token, err := a.oidcVerifier.Verify(ctx, bearerToken)
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "invalid jwt: %v", err)
		}

		claim, err := a.getClaims(token)
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "could not parse claim: %v", err)
		}

		userInDB, err := getUserFromDatabase(ctx, configDatabase, claim.Email)
		if err != nil {
			return nil, err
		}

		newContext := context.WithValue(ctx, domain.UserKey, userInDB)

		newContextWithUserAgent := context.WithValue(newContext, domain.UserAgentKey, userAgent)

		return handler(newContextWithUserAgent, req)
	}
}

func (a *Authenticator) StreamServerInterceptor(_ database.ConfigDatabase, _ func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return fmt.Errorf("StreamServerInterceptor unimplemented")
	}
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}
