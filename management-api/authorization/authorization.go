package authorization

import (
	"net/http"
)

// Authorizer defines the contract to authorize a request
type Authorizer interface {
	Authorize(r *http.Request) bool
}

// Authorization provides a middleware for an Authorizer
type Authorization struct {
	authorizer Authorizer
}

// NewAuthorization creates an Authorization
func NewAuthorization(authorizer Authorizer) *Authorization {
	return &Authorization{authorizer}
}

// Middleware returns StatusForbidden for Requests that fail to meet the conditions of the Authorizer
func (a Authorization) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.authorizer.Authorize(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})
}
