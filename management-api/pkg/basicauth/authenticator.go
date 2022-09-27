// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/api"
	"go.nlx.io/nlx/management-api/pkg/auditlog"
	"go.nlx.io/nlx/management-api/pkg/database"
)

type Authenticator struct {
	configDatabase database.ConfigDatabase
	logger         *zap.Logger
	auditLogger    auditlog.Logger
}

func NewAuthenticator(configDatabase database.ConfigDatabase, auditLogger auditlog.Logger, logger *zap.Logger) api.Authenticator {
	return &Authenticator{
		configDatabase: configDatabase,
		auditLogger:    auditLogger,
		logger:         logger,
	}
}

func (a *Authenticator) MountRoutes(r chi.Router) {
	routes := chi.NewRouter()
	routes.Get("/", a.root)
	routes.Get("/me", a.me)
	routes.Post("/login", a.login)

	r.Mount("/basic-auth", routes)
}

func (a *Authenticator) OnlyAuthenticated(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		credentialsMatch, err := a.configDatabase.VerifyUserCredentials(context.Background(), username, password)
		if err != nil {
			a.logger.Info("Unable to verify user credentials", zap.String("username", username))
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		if !credentialsMatch {
			a.logger.Info("Unable to verify user credentials", zap.Error(err))
			http.Error(w, "unauthorized request", http.StatusUnauthorized)
			return
		}

		r.Header.Add("Authorization", "Basic "+encodeBasicAuth(username, password))

		h.ServeHTTP(w, r)
	})
}

// NOTE: This endpoint is used clientside to check whether to use basic auth or OIDC
func (a *Authenticator) root(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (a *Authenticator) me(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	valid, err := a.configDatabase.VerifyUserCredentials(context.Background(), username, password)
	if err != nil {
		a.logger.Error("Unable to verify user credentials", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if !valid {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	user, err := a.configDatabase.GetUser(context.Background(), username)
	if err != nil {
		a.logger.Error("Unable to retrieve user from database", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	type User struct {
		ID         string `json:"id"`
		FullName   string `json:"fullName"`
		Email      string `json:"email"`
		PictureURL string `json:"pictureUrl"`
	}

	render.JSON(w, r, &User{
		ID:         fmt.Sprintf("%v", user.ID),
		FullName:   "",
		Email:      user.Email,
		PictureURL: "",
	})
}

func (a *Authenticator) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if len(email) < 1 || len(password) < 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	credentialsMatch, err := a.configDatabase.VerifyUserCredentials(context.Background(), email, password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if !credentialsMatch {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = a.auditLogger.LoginSuccess(r.Context(), email, r.Header.Get("User-Agent"))
	if err != nil {
		a.logger.Error("error writing to audit log", zap.Error(err))
	}

	w.WriteHeader(http.StatusNoContent)
}

func (a *Authenticator) UnaryServerInterceptor(configDatabase database.ConfigDatabase, getUserFromDatabase func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not extract metadata from context")
		}

		authHeader := md.Get("Authorization")

		// Outway and Inway do not send authorization header
		if len(authHeader) == 0 {
			return handler(ctx, req)
		}

		var userAgent string
		if len(md.Get("grpcgateway-user-agent")) > 0 {
			userAgent = md.Get("grpcgateway-user-agent")[0]
		} else if len(md.Get("user-agent")) > 0 {
			userAgent = md.Get("user-agent")[0]
		}

		username, password, ok := getBasicAuth(authHeader[0])
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "invalid or empty Authorization header")
		}

		valid, err := a.configDatabase.VerifyUserCredentials(context.Background(), username, password)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not verify user credentials: %v", err)
		}

		if !valid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
		}

		userInDB, err := getUserFromDatabase(ctx, configDatabase, username)
		if err != nil {
			return nil, err
		}

		contextWithUser := context.WithValue(ctx, domain.UserKey, userInDB)
		contextWithUserAgent := context.WithValue(contextWithUser, domain.UserAgentKey, userAgent)

		return handler(contextWithUserAgent, req)
	}
}

func (a *Authenticator) StreamServerInterceptor(_ database.ConfigDatabase, _ func(ctx context.Context, configDatabase database.ConfigDatabase, email string) (*domain.User, error)) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return fmt.Errorf("StreamServerInterceptor unimplemented")
	}
}

func getBasicAuth(auth string) (username, password string, ok bool) {
	if auth == "" {
		return "", "", false
	}

	return parseBasicAuth(auth)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return "", "", false
	}

	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return "", "", false
	}

	cs := string(c)

	username, password, ok = strings.Cut(cs, ":")
	if !ok {
		return "", "", false
	}

	return username, password, true
}

func encodeBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
