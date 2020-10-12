// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package irma

import (
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// ResultCode is used in VerificationResultClaims to indicate the status/result of a verification session.
type ResultCode string

const (
	// ResultCodeWaiting indicates that the client has not yet posted a proof
	ResultCodeWaiting ResultCode = `WAITING`
	// ResultCodeExpired indicates one or more of the proofs came from an expired credential.
	ResultCodeExpired ResultCode = `EXPIRED`
	// ResultCodeInvalid indicates the proofs were invalid.
	ResultCodeInvalid ResultCode = `INVALID`
	// ResultCodeMissingAttributes indicates not all required attributes were disclosed in the proofs
	ResultCodeMissingAttributes ResultCode = `MISSING_ATTRIBUTES`
	// ResultCodeValid indicates the proofs were valid.
	ResultCodeValid ResultCode = `VALID`
)

type VerificationResultClaims struct {
	jwt.StandardClaims
	Result     ResultCode        `json:"result"`
	Attributes map[string]string `json:"attributes"`
}

type JWTHandler interface {
	VerifyIRMAVerificationResult(jwtBytes []byte, rsaVerifyPublicKey *rsa.PublicKey) (*jwt.Token, *VerificationResultClaims, error)
}

type JWTGenerator struct {
}

func NewJWTGenerator() *JWTGenerator {
	return &JWTGenerator{}
}

func jwtVerifyKeyFunc(rsaVerifyPublicKey *rsa.PublicKey) jwt.Keyfunc {
	return func(*jwt.Token) (interface{}, error) {
		return rsaVerifyPublicKey, nil
	}
}

// VerifyIRMAVerificationResult unpacks and validates IRMA Vericicaiotn Result JWT and returns the token and claims or an error.
func (j *JWTGenerator) VerifyIRMAVerificationResult(jwtBytes []byte, rsaVerifyPublicKey *rsa.PublicKey) (*jwt.Token, *VerificationResultClaims, error) {
	claims := &VerificationResultClaims{}

	token, err := jwt.ParseWithClaims(string(jwtBytes), claims, jwtVerifyKeyFunc(rsaVerifyPublicKey))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse JWT received from irma-api-server")
	}

	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, nil, errors.Errorf("JWT received from irma-api-server uses unexpected signing method %v", token.Header["alg"])
	}

	if !token.Valid {
		return nil, nil, errors.New("JWT received from irma-api-server is invalid")
	}

	return token, claims, nil
}
