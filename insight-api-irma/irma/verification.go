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

// VerificationDiscloseSession holds details for a disclose session as opened on the irma-api-server
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

// StartVerification sends a verification request and returns the verification disclose session object.
func (c *Client) StartVerification(request *DiscloseRequest) (*VerificationDiscloseSession, error) {
	// prepare and sign jwt
	claims := c.newVerificationRequestClaims(request)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedJWT, err := token.SignedString(c.RSASignPrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign jwt")
	}
	c.logger.Debug("sending verification request to irma-api-server", zap.String("jwt", signedJWT))

	// make http request
	resp, err := http.Post(c.IRMAEndpointURL.String()+"/api/v2/verification", "text/plain", bytes.NewBufferString(signedJWT))
	if err != nil {
		return nil, errors.Wrap(err, "failed to do http request to irma-api-server")
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

// PollVerification polls the irma-api-server for the status or result of a verification session.
func (c *Client) PollVerification(verificationID string) (string, *VerificationResultClaims, error) {
	// make http request
	resp, err := http.Get(c.IRMAEndpointURL.String() + "/api/v2/verification/" + verificationID)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to do http request to irma-api-server")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		return "", nil, errors.New("failed to start verification, http status " + resp.Status + " and data " + string(data))
	}

	jwtResult, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to read response body from irma-api-server")
	}

	// unpack response
	claims := &VerificationResultClaims{}
	token, err := jwt.ParseWithClaims(string(jwtResult), claims, c.jwtVerifyKeyFunc)
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to parse JWT received from irma-api-server")
	}
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return "", nil, errors.Errorf("JWT received from irma-api-server uses unexpected signing method %v", token.Header["alg"])
	}
	if !token.Valid {
		return "", nil, errors.New("JWT received from irma-api-server is invalid")
	}

	return token.Raw, claims, nil
}
