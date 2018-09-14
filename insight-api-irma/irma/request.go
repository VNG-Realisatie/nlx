package irma

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

type VerificationRequestClaims struct {
	jwt.StandardClaims
	SPRequest SPRequestClaims `json:"sprequest"`
}

type SPRequestClaims struct {
	Data        string           `json:"data"`
	Validity    uint64           `json:"validity"`
	Timeout     uint64           `json:"timeout"`
	CallbackURL string           `json:"callbackUrl"`
	Request     *DiscloseRequest `json:"request"`
}

// VerificationDiscloseSession holds details for a disclose session as opened on the irma api server
type VerificationDiscloseSession struct {
	IRMAQR     string `json:"irmaqr"`
	U          string `json:"u"`
	Version    string `json:"v"`
	VersionMax string `json:"vmax"`
}

// newVerificationRequestClaims creates a new VerificationClaims object for given DiscloseRequest
func (c *Client) newVerificationRequestClaims(request *DiscloseRequest) *VerificationRequestClaims {
	claims := &VerificationRequestClaims{
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
func (c *Client) SendVerificationRequest(request *DiscloseRequest) (*VerificationDiscloseSession, error) {
	// prepare and sign jwt
	claims := c.newVerificationRequestClaims(request)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedJWT, err := token.SignedString(c.RSASigningKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign jwt")
	}
	c.logger.Debug("sending verification request to irma api server", zap.String("jwt", signedJWT))

	// make http request
	resp, err := http.Post(c.IRMAEndpointURL.String()+"/api/v2/verification", "text/plain", bytes.NewBufferString(signedJWT))
	if err != nil {
		return nil, errors.Wrap(err, "failed to do http request to irma api server")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.New("failed to start verification, http status " + resp.Status + " and data " + string(data))
	}

	// unpack response
	session := &VerificationDiscloseSession{}
	err = json.NewDecoder(resp.Body).Decode(session)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall response")
	}
	return session, nil
}
