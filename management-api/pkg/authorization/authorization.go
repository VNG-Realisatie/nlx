// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package authorization

import "go.nlx.io/nlx/management-api/pkg/permissions"

func IsAuthorized(permission permissions.Permission, authorizedPermissions map[permissions.Permission]bool) bool {
	return authorizedPermissions[permission]
}
