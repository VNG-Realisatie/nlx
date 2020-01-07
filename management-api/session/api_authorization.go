package session

import (
	"net/http"
	"regexp"

	"go.nlx.io/nlx/management-api/models"
)

type permission struct {
	Method       string
	PathsAllowed *regexp.Regexp
	RolesAllowed []string
}

var permissions = [...]permission{
	{PathsAllowed: regexp.MustCompile("^/api/auth/.*$")},
	{Method: "GET", RolesAllowed: []string{"admin", "readonly"}}, // Role must be readonly or admin
	{Method: "POST", RolesAllowed: []string{"admin"}},            // Role must be admin (using a case insentive match)
	{Method: "PUT", RolesAllowed: []string{"admin"}},             // Role must be admin (using a case insentive match)
}

// Authorizer for a Session
type Authorizer struct {
}

// NewAuthorizer creates a new Authorizer
func NewAuthorizer() *Authorizer {
	return &Authorizer{}
}

// Authorize allow access to based on static conditions using the session from the context
func (a Authorizer) Authorize(r *http.Request) bool {
	session := getSession(r)

	for _, p := range permissions {
		if p.Method == "" || p.Method == r.Method {
			if p.RolesAllowed != nil && session != nil {
				if account, _ := session.Account(); account != nil {
					b := p.isRoleAllowedFor(account)
					if b {
						return true
					}
				}
			} else if p.PathsAllowed != nil && p.PathsAllowed.MatchString(r.URL.Path) {
				return true
			}
		}
	}

	return false
}

func (p permission) isRoleAllowedFor(account *models.Account) bool {
	for _, role := range p.RolesAllowed {
		if role == account.Role {
			return true
		}
	}

	return false
}
