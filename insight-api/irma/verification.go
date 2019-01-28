package irma

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type VerificationRequestClaims struct {
	jwt.StandardClaims
	SPRequest SPRequestClaims `json:"sprequest"`
}

// SPRequestClaims contains the fields for a verification request
type SPRequestClaims struct {
	Data        string           `json:"data"`
	Validity    uint64           `json:"validity"`
	Timeout     uint64           `json:"timeout"`
	CallbackURL string           `json:"callbackUrl"`
	Request     *DiscloseRequest `json:"request"`
}

// DiscloseRequest contains the data for a disclose request
type DiscloseRequest struct {
	Content []DiscloseRequestContent `json:"content"`
}

// DiscloseRequestContent contains information about a required attribute(set) in a disclose request
type DiscloseRequestContent struct {
	Label      string      `json:"label"`
	Attributes []Attribute `json:"attributes"`
}

// newVerificationRequestClaims creates a new VerificationClaims object for given DiscloseRequest
func newVerificationRequestClaims(request *DiscloseRequest, serviceProviderName string) *VerificationRequestClaims {
	claims := &VerificationRequestClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:   serviceProviderName,
			IssuedAt: time.Now().Unix(),
			Subject:  "verification_request",
		},
		SPRequest: SPRequestClaims{
			Data:     "",
			Validity: 3600,
			Timeout:  60,
			Request:  request,
		},
	}
	return claims
}

// {
//     "exp": 1448636691,
//     "sub": "disclosure_result",
//     "jti": "foobar",
//     "attributes": {
//         "irma-demo.MijnOverheid.ageLower.over18": "yes",
//         "irma-demo.IRMATube.member": "present",
//     },
//     "iat": 1448636631,
//     "status": "VALID"
// }

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

func GenerateAndSignJWT(request *DiscloseRequest, serviceProviderName string, rsaSignPrivateKey *rsa.PrivateKey) (string, error) {
	claims := newVerificationRequestClaims(request, serviceProviderName)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedJWT, err := token.SignedString(rsaSignPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign jwt")
	}
	return signedJWT, nil
}

func jwtVerifyKeyFunc(rsaVerifyPublicKey *rsa.PublicKey) jwt.Keyfunc {
	return func(*jwt.Token) (interface{}, error) {
		return rsaVerifyPublicKey, nil
	}
}

// VerifyIRMAVerificationResult unpacks and validates IRMA Vericicaiotn Result JWT and returns the token and claims or an error.
func VerifyIRMAVerificationResult(jwtBytes []byte, rsaVerifyPublicKey *rsa.PublicKey) (*jwt.Token, *VerificationResultClaims, error) {
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
