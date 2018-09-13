package irma

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// "iss": "Service provider name",
// "kid": "service_provider_name",
// "sub": "verification_request",
// "iat": 1453377600,
// "sprequest": {
//     "data": "...",
//     "validity": 60,
//     "timeout": 60,
//     "callbackUrl": "...",
//     "request": "..."
// }

type VerificationClaims struct {
	jwt.StandardClaims
	SPRequest SPRequestClaims `json:"sprequest"`
}

type SPRequestClaims struct {
	Data        string           `json:"data"`
	Validity    uint64           `json:"validity"`
	Timeout     uint64           `json:"timeout"`
	CallbackURL url.URL          `json:"callbackUrl"`
	Request     *DiscloseRequest `json:"request"`
}

type verificationResult struct {
	U string `json:"u"`
}

// newVerificationRequestClaims creates a new VerificationClaims object for given DiscloseRequest
func (c *Client) newVerificationRequestClaims(request *DiscloseRequest) *VerificationClaims {
	claims := &VerificationClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:   c.ServiceProviderName,
			IssuedAt: time.Now().Unix(),
			Subject:  "verification_request",
		},
		SPRequest: SPRequestClaims{
			Data:     "",
			Validity: 60,
			Timeout:  60,
			Request:  request,
		},
	}
	return claims
}

// SendVerificationRequest sends a verification request and returns the verificationID
func (c *Client) SendVerificationRequest(request *DiscloseRequest) (string, error) {
	// prepare and sign jwt
	claims := c.newVerificationRequestClaims(request)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(c.RSASigningKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign jwt")
	}

	// make http request
	resp, err := http.Post(c.IRMAEndpointURL.String()+"/api/v2/verification", "", bytes.NewBufferString(signedToken))
	if err != nil {
		return "", errors.Wrap(err, "failed to do http request to irma api server")
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to start verification, http status " + resp.Status)
	}
	defer resp.Body.Close()

	// unpack response
	result := verificationResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshall response")
	}
	return result.U, nil
}
