// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package httperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Code int
type Source int
type Location int

type NLXNetworkError struct {
	Message  string   `json:"message"`
	Source   Source   `json:"source"`
	Location Location `json:"location,omitempty"`
	Code     Code     `json:"code"`
}

const StatusNLXNetworkError = 540

func WriteError(w http.ResponseWriter, source Source, location Location, code Code, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(StatusNLXNetworkError)

	networkErr, err := json.Marshal(&NLXNetworkError{
		Source:   source,
		Location: location,
		Code:     code,
		Message:  message,
	})
	if err != nil {
		fmt.Fprintf(w, `{"source":%q,"code":%q,"message":"error while marshaling json error"}`, source.String(), ServerError.String())
		return
	}

	_, _ = w.Write(networkErr)
}

// Error codes
// The comment after the error code is the string representation of the error code
const (
	// Empty path when calling inway, must at least contain the service name.
	EmptyPath Code = iota + 1 // EMPTY_PATH

	// The requested service does not exist.
	ServiceDoesNotExist // SERVICE_DOES_NOT_EXIST

	// There was an error while executing the plugin chain
	ErrorExecutingPluginChain // ERROR_EXECUTING_PLUGIN_CHAIN

	// The service is either unreachable or down.
	ServiceUnreachable // SERVICE_UNREACHABLE

	// Missing peer certificate in connection to inway
	MissingPeerCertificate // MISSING_PEER_CERTIFICATE

	// Invalid certificate provided in connection to inway.
	// Certificate must contain organization name, subject serial number and a organization issuer
	InvalidCertificate // INVALID_CERTIFICATE

	// Access denied, no valid access grant was found,
	// request access to this service via the Management
	AccessDenied // ACCESS_DENIED

	// There was an error while authorizing the request via the authorization server
	ErrorWhileAuthorizingRequest // ERROR_WHILE_AUTHORIZING_REQUEST

	// The authorization server denied the request
	Unauthorized // UNAUTHORIZED

	// Unable to verify claim for delegation, claim can not be parsed
	UnableToVerifyClaim // UNABLE_TO_VERIFY_CLAIM

	// The requesting organization is not the organization found in the order
	RequestingOrganizationIsNotDelegatee // REQUESTING_ORGANIZATION_IS_NOT_DELEGATEE

	// The delegator of this order does not have access to the service
	DelegatorDoesNotHaveAccessToService // DELEGATOR_DOES_NOT_HAVE_ACCESS_TO_SERVICE

	// Missing log record ID, the header 'X-NLX-Logrecord-Id' must be set with an unique ID for this request
	MissingLogRecordID // MISSING_LOG_RECORD_ID

	// General server error, see message for more information
	ServerError // SERVER_ERROR

	// Outway is called with an invalid URL, see message for more information
	InvalidURL // INVALID_URL

	// Proxy mode is disabled, enable it by setting the 'use-as-http-proxy' flag to resolve
	ProxyModeDisabled // PROXY_MODE_DISABLED

	// Outway got called with an invalid method, CONNECT method is not supported
	UnsupportedMethod // UNSUPPORTED_METHOD

	// Unable to parse delegation metadata,
	// check if 'X-NLX-Request-Delegator' and 'X-NLX-Request-Order-Reference' headers are set correctly.
	UnableToParseDelegationMetadata // UNABLE_TO_PARSE_DELEGATION_METADATA

	// Unable to setup management client to retrieve claim for order
	UnableToSetupManagementClient // UNABLE_TO_SETUP_MANAGEMENT_CLIENT

	// Order was not found
	OrderNotFound // ORDER_NOT_FOUND

	// The used order does not exist for your organization
	OrderDoesNotExistForYourOrganization // ORDER_DOES_NOT_EXIST_FOR_YOUR_ORGANIZATION

	// The order is revoked by the delegator
	OrderRevoked // ORDER_REVOKED

	// The order is expired
	OrderExpired // ORDER_EXPIRED

	// The order does not contain the service that was tried to access
	OrderDoesNotContainService // ORDER_DOES_NOT_CONTAIN_SERVICE

	// Unable to request claim from delegator
	UnableToRequestClaim // UNABLE_TO_REQUEST_CLAIM

	// Received an invalid claim from the delegator
	ReceivedInvalidClaim // RECEIVED_INVALID_CLAIM

	// Invalid data subject header, 'X-NLX-Request-Data-Subject' contains invalid data.
	// Must be in 'key=value' format.
	InvalidDataSubjectHeader // INVALID_DATA_SUBJECT_HEADER
)

// Source
// The comment after the error code is the string representation of the error code
const (
	// The error originated from the Inway
	Inway Source = iota + 1 // inway

	// The error originated from the Outway
	Outway // outway
)

// Location
// Check https://docs.nlx.io/support/common-errors for a graphical overview
// The comment after the error code is the string representation of the error code
const (
	// The error happened between the Inway and the API
	A1 Location = iota + 1

	// The error happened between the Inway and the Authorization server
	IAS1

	// The error happened between the Client and the Outway
	C1

	// The error happened between the Directory Monitor and the Inway.
	// This means that the Inway couldn't be reached from the Directory Monitor and thus Outways cannot reach your Inway
	M1

	// The error happened between the Outway and Inway
	O1

	// The error happened between the Outway and the Authorization server
	OAS1
)

func (l Location) GoString() string {
	return l.String()
}

func (l Code) GoString() string {
	return l.String()
}

func (l Source) GoString() string {
	return l.String()
}
