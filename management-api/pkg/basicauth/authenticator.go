// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package basicauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"go.uber.org/zap"

	"go.nlx.io/nlx/management-api/pkg/api"
	"go.nlx.io/nlx/management-api/pkg/database"
)

type Authenticator struct {
	configDatabase database.ConfigDatabase
	logger         *zap.Logger
}

func NewAuthenticator(configDatabase database.ConfigDatabase, logger *zap.Logger) api.Authenticator {
	return &Authenticator{
		configDatabase: configDatabase,
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

		r.Header.Add("username", username)

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

	w.WriteHeader(http.StatusNoContent)
}
