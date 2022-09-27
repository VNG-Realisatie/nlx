// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package domain

import "go.nlx.io/nlx/management-api/pkg/permissions"

type User struct {
	Email       string
	Permissions map[permissions.Permission]bool
}

type UserAgentContextKey string
type UserContextKey string

var UserAgentKey UserAgentContextKey = "userAgentKey"
var UserKey UserContextKey = "userKey"
