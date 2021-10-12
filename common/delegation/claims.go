// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package delegation

import "github.com/golang-jwt/jwt"

type Service struct {
	Service                  string `json:"service"`
	OrganizationSerialNumber string `json:"organization_serial_number"`
}

type JWTClaims struct {
	jwt.StandardClaims
	Delegatee      string    `json:"delegatee"`
	OrderReference string    `json:"orderReference"`
	Services       []Service `json:"services"`
}
