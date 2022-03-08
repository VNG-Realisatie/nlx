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
	Delegatee                     string       `json:"delegatee"`
	DelegateePublicKeyFingerprint string       `json:"delegatee_public_key_fingerprint"`
	OrderReference                string       `json:"orderReference"`
	AccessProof                   *AccessProof `json:"accessProof"`
}

func (j *JWTClaims) IsValidFor(serviceName, organizationSerialNumber, publicKeyFingerprint string) bool {
	if j.AccessProof.ServiceName == serviceName && j.AccessProof.OrganizationSerialNumber == organizationSerialNumber && publicKeyFingerprint == j.AccessProof.PublicKeyFingerprint {
		return true
	}

	return false
}
