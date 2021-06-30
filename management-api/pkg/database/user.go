package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type RoleNotFoundError struct {
	RoleName string
}

func (err *RoleNotFoundError) Error() string {
	return fmt.Sprintf("role '%s' not found", err.RoleName)
}

const AdminRole = "admin"

type User struct {
	ID        uint
	Email     string
	Password  string
	Roles     []Role `gorm:"many2many:nlx_management.users_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*User) TableName() string {
	return "nlx_management.users"
}

func (user *User) HasRole(code string) bool {
	for _, role := range user.Roles {
		if role.Code == code {
			return true
		}
	}

	return false
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (db *PostgresConfigDatabase) GetUser(ctx context.Context, email string) (*User, error) {
	user := &User{}

	if err := db.
		WithContext(ctx).
		Where("email = ?", email).
		Preload("Roles").
		First(user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (db *PostgresConfigDatabase) VerifyUserCredentials(ctx context.Context, email, password string) (bool, error) {
	user := &User{}
	if err := db.
		WithContext(ctx).
		Where("email = ?", email).
		Preload("Roles").
		First(user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}

		return false, err
	}

	match := checkPasswordHash(password, user.Password)

	return match, nil
}

func (db *PostgresConfigDatabase) CreateUser(ctx context.Context, email, password string, roleNames []string) (*User, error) {
	tx := db.DB.Begin()
	defer tx.Rollback()

	dbWithTx := &PostgresConfigDatabase{
		DB: tx,
	}

	roles, err := getRoleRecords(ctx, dbWithTx, roleNames)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email: email,
	}

	if len(password) > 0 {
		hashedPassword, hashErr := hashPassword(password)
		if hashErr != nil {
			return nil, fmt.Errorf("failed to hash password: %v", hashErr)
		}

		user.Password = hashedPassword
	}

	var count int64

	if err := dbWithTx.
		WithContext(ctx).
		Model(User{}).
		Where("email = ?", email).
		Count(&count).
		Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, ErrUserAlreadyExists
	}

	if err := dbWithTx.
		WithContext(ctx).
		Create(user).
		Error; err != nil {
		return nil, err
	}

	err = dbWithTx.
		WithContext(ctx).
		Model(user).
		Association("Roles").
		Append(roles)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

func getRoleRecords(ctx context.Context, dbWithTx *PostgresConfigDatabase, names []string) ([]Role, error) {
	roles := []Role{}

	for _, name := range names {
		role := &Role{}

		if err := dbWithTx.
			WithContext(ctx).
			Where("code = ?", name).
			First(role).
			Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &RoleNotFoundError{
					RoleName: name,
				}
			}

			return nil, err
		}

		roles = append(roles, *role)
	}

	return roles, nil
}
