// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

//go:build integration

package database_test

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/permissions"
)

func TestHasPermissions(t *testing.T) {
	t.Parallel()

	setup(t)

	configDb, close := newConfigDatabase(t, t.Name(), false)
	defer close()

	permissionsFromDB, err := configDb.ListPermissions(context.Background())
	assert.Nil(t, err)

	permissionsFromDBMap := make(map[string]bool)
	for _, permission := range permissionsFromDB {
		_, err := permissions.PermissionString(permission.Code)
		if err != nil {
			t.Errorf("invalid permission in database: %q", permission.Code)
		}

		permissionsFromDBMap[permission.Code] = true
	}

	missingPermissions := []string{}
	for _, permission := range permissions.PermissionValues() {
		_, ok := permissionsFromDBMap[permission.String()]
		if !ok {
			missingPermissions = append(missingPermissions, permission.String())
		}
	}

	if len(missingPermissions) > 0 {
		insertMissingPermissionsQuery := "INSERT INTO nlx_management.permissions (code) VALUES\n"
		for i, p := range missingPermissions {
			insertMissingPermissionsQuery = fmt.Sprintf("%s\t('%s')", insertMissingPermissionsQuery, p)

			if i == len(missingPermissions)-1 {
				insertMissingPermissionsQuery = fmt.Sprintf("%s;", insertMissingPermissionsQuery)
			} else {
				insertMissingPermissionsQuery = fmt.Sprintf("%s,\n", insertMissingPermissionsQuery)
			}
		}

		t.Errorf("missing permissions in database, insert with following query:\n%s", insertMissingPermissionsQuery)
	}
}
