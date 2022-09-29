// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package domain

import "go.nlx.io/nlx/management-api/pkg/permissions"

type User struct {
	ID       uint
	Email    string
	Password string
	Roles    []*Role
}

type Role struct {
	Code        string
	Permissions []permissions.Permission
}

type UserAgentContextKey string
type UserContextKey string

var UserAgentKey UserAgentContextKey = "userAgentKey"
var UserKey UserContextKey = "userKey"
