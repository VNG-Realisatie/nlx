// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package domain

import "go.nlx.io/nlx/management-api/pkg/permissions"

type User struct {
	Email       string
	UserAgent   string
	Permissions map[permissions.Permission]bool
}

type UserContextKey string

var UserKey UserContextKey = "userKey"
