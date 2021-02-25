// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package oidc

import (
	"errors"
	"time"
)

type Claims struct {
	Subject   string `json:"sub"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	Issuer    string `json:"iss"`
	Audience  string `json:"aud"`
	ExpiresAt int64  `json:"exp,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
}

func (claims *Claims) Verify() error {
	now := time.Now().Unix()

	if claims.ExpiresAt > 0 {
		if now >= claims.ExpiresAt {
			return errors.New("JWT is expired")
		}
	}

	if claims.IssuedAt > 0 {
		if now < claims.IssuedAt {
			return errors.New("cannot use JWT before it has been issued")
		}
	}

	if claims.NotBefore > 0 {
		if now <= claims.NotBefore {
			return errors.New("cannot use JWT before it's valid")
		}
	}

	return nil
}

func (claims *Claims) User() *User {
	return &User{
		ID:         claims.Subject,
		FullName:   claims.Name,
		Email:      claims.Email,
		PictureURL: claims.Picture,
	}
}

type User struct {
	ID         string `json:"id"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	PictureURL string `json:"pictureUrl"`
}
