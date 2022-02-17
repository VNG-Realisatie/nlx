// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package delegation

import (
	"github.com/golang-jwt/jwt/v4"
)

type AccessProof struct {
	ServiceName              string `json:"service_name"`
	OrganizationSerialNumber string `json:"organization_serial_number"`
	PublicKeyFingerprint     string `json:"public_key_fingerprint"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Delegatee      string         `json:"delegatee"`
	OrderReference string         `json:"orderReference"`
	AccessProofs   []*AccessProof `json:"accessProofs"`
}

func (j *JWTClaims) IsValidFor(serviceName, organizationSerialNumber, publicKeyFingerprint string) bool {
	for _, accessProof := range j.AccessProofs {
		return accessProof.ServiceName == serviceName && accessProof.OrganizationSerialNumber == organizationSerialNumber && publicKeyFingerprint == accessProof.PublicKeyFingerprint
	}

	return false
}
