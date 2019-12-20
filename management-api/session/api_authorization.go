package session

import (
	"net/http"
	"regexp"
)

var permissions = [...]struct {
	Method      string
	PathPattern *regexp.Regexp
	RolePattern *regexp.Regexp
}{
	{PathPattern: regexp.MustCompile("^/api/auth/.*$")},
	{Method: "GET", RolePattern: regexp.MustCompile("^(?i:readonly|admin)$")}, // Role must be readonly or admin
	{Method: "POST", RolePattern: regexp.MustCompile("^(?i:admin)$")},         // Role must be admin (using a case insentive match)
	{Method: "PUT", RolePattern: regexp.MustCompile("^(?i:admin)$")},          // Role must be admin (using a case insentive match)
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
			if p.RolePattern != nil && session != nil {
				if account, _ := session.Account(); account != nil {
					if p.RolePattern.MatchString(account.Role) {
						return true
					}
				}
			} else if p.PathPattern != nil && p.PathPattern.MatchString(r.URL.Path) {
				return true
			}
		}
	}

	return false
}
