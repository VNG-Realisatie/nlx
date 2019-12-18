package session

import (
	"net/http"
	"regexp"
)

var permissions = [...]struct {
	Method      string
	RolePattern *regexp.Regexp
}{
	{Method: "GET", RolePattern: regexp.MustCompile("^(?i:readonly|admin)$")}, // Role must be readonly or admin
	{Method: "POST", RolePattern: regexp.MustCompile("^(?i:admin)$")},         // Role must be admin (using a case insentive match)
	{Method: "PUT", RolePattern: regexp.MustCompile("^(?i:admin)$")},          // Role must be admin (using a case insentive match)
}

type Authorizer struct{}

// Authorize allow access to based on static conditions
func (sa Authorizer) Authorize(r *http.Request) bool {
	if session := getSession(r); session != nil {
		account, err := session.Account()
		if err == nil {
			for _, p := range permissions {
				if p.Method == r.Method && p.RolePattern.MatchString(account.Role) {
					return true
				}
			}
		}
	}

	return false
}
