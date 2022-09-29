// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.nlx.io/nlx/management-api/adapters/storage/postgres/queries"
	"go.nlx.io/nlx/management-api/domain"
	"go.nlx.io/nlx/management-api/pkg/permissions"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (db *PostgresConfigDatabase) GetUser(ctx context.Context, email string) (*domain.User, error) {
	user, err := db.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	roles, err := db.queries.ListRolesForUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	roleModels := make([]*domain.Role, len(roles))

	for i, role := range roles {
		permissionsForRole, err := db.queries.ListPermissionsForRole(ctx, role)
		if err != nil {
			return nil, err
		}

		permissionModels := make([]permissions.Permission, len(permissionsForRole))

		for j, permission := range permissionsForRole {
			p, err := permissions.PermissionString(permission)
			if err != nil {
				return nil, fmt.Errorf("invalid permission %q", permission)
			}

			permissionModels[j] = p
		}

		roleModels[i] = &domain.Role{
			Code:        role,
			Permissions: permissionModels,
		}
	}

	result := &domain.User{
		ID:    uint(user.ID),
		Email: user.Email,
		Roles: roleModels,
	}

	return result, nil
}

func (db *PostgresConfigDatabase) VerifyUserCredentials(ctx context.Context, email, password string) (bool, error) {
	user, err := db.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ErrNotFound
		}

		return false, err
	}

	result := false
	if user.Password.Valid {
		result = checkPasswordHash(password, user.Password.String)
	}

	return result, nil
}

// nolint:gocyclo // unable to reduce complexity
func (db *PostgresConfigDatabase) CreateUser(ctx context.Context, email, password string, roleNames []string) (id uint, anError error) {
	tx, err := db.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		err = tx.Rollback()
		if err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}

			fmt.Printf("cannot rollback database transaction for creating a user: %e", err)
		}
	}()

	qtx := db.queries.WithTx(tx)

	hashedPassword, hashErr := hashPassword(password)
	if hashErr != nil {
		return 0, fmt.Errorf("failed to hash password: %v", hashErr)
	}

	now := time.Now()

	userID, err := qtx.CreateUser(ctx, &queries.CreateUserParams{
		Email: email,
		Password: sql.NullString{
			Valid:  true,
			String: hashedPassword,
		},
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		if strings.Contains(err.Error(), "value violates unique constraint \"users_email_key\"") {
			return 0, ErrUserAlreadyExists
		}

		return 0, err
	}

	err = assignUserRoles(ctx, qtx, uint(userID), roleNames)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}

func assignUserRoles(ctx context.Context, qtx *queries.Queries, userID uint, roleNames []string) error {
	now := time.Now()

	for _, role := range roleNames {
		err := qtx.CreateUserRoles(ctx, &queries.CreateUserRolesParams{
			UserID:    int32(userID),
			RoleCode:  role,
			CreatedAt: now,
			UpdatedAt: now,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
