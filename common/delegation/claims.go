// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package delegation

import "github.com/form3tech-oss/jwt-go"

type JWTClaims struct {
	jwt.StandardClaims
	Organization   string `json:"organization"`
	OrderReference string `json:"order_reference"`
}
