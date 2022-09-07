// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package httperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Code string
type Source string
type Location string

type NLXNetworkError struct {
	Message  string   `json:"message"`
	Source   Source   `json:"source"`
	Location Location `json:"location,omitempty"`
	Code     Code     `json:"code"`
}

const StatusNLXNetworkError = 540

func WriteError(w http.ResponseWriter, nlxErr *NLXNetworkError) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(StatusNLXNetworkError)

	networkErr, err := json.Marshal(nlxErr)
	if err != nil {
		fmt.Fprintf(w, `{"source":%q,"code":%q,"message":"error while marshaling json error"}`, nlxErr.Source.String(), ServerErrorErr.String())
		return
	}

	_, _ = w.Write(networkErr)
}

// The requested service does not exist.
const ServiceDoesNotExistErr Code = "SERVICE_DOES_NOT_EXIST"

func ServiceDoesNotExist(serviceName string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ServiceDoesNotExistErr,
		Message: fmt.Sprintf("no endpoint for service '%s'", serviceName),
	}
}

// Empty path when calling inway, must at least contain the service name.
const EmptyPathErr Code = "EMPTY_PATH"

func EmptyPath() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    EmptyPathErr,
		Message: "path cannot be empty, must at least contain the service name.",
	}
}

// There was an error while executing the plugin chain
const ErrorExecutingPluginChainErr Code = "ERROR_EXECUTING_PLUGIN_CHAIN"

func ErrorExecutingPluginChain() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ErrorExecutingPluginChainErr,
		Message: "error executing plugin chain",
	}
}

// The service is either unreachable or down.
const ServiceUnreachableErr Code = "SERVICE_UNREACHABLE"

func ServiceUnreachable(serviceURL string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ServiceUnreachableErr,
		Message: fmt.Sprintf("failed API request to %s try again later. service api down/unreachable. check error at https://docs.nlx.io/support/common-errors/", serviceURL),
	}
}

// Missing peer certificate in connection to inway
const MissingPeerCertificateErr Code = "MISSING_PEER_CERTIFICATE"

func MissingPeerCertificate() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    MissingPeerCertificateErr,
		Message: "invalid connection: missing peer certificates",
	}
}

// Invalid certificate provided in connection to inway.
// Certificate must contain organization name, subject serial number and a organization issuer
const InvalidCertificateErr Code = "INVALID_CERTIFICATE"

func InvalidCertificate(msg string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    InvalidCertificateErr,
		Message: msg,
	}
}

// Access denied, no valid access grant was found,
// request access to this service via the Management
const AccessDeniedErr Code = "ACCESS_DENIED"

func AccessDenied(orgSerialNumber, orgPublicKeyFingerprint string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    AccessDeniedErr,
		Message: fmt.Sprintf(`permission denied, organization %q or public key fingerprint %q is not allowed access.`, orgSerialNumber, orgPublicKeyFingerprint),
	}
}

// There was an error while authorizing the request via the authorization server
const ErrorWhileAuthorizingRequestErr Code = "ERROR_WHILE_AUTHORIZING_REQUEST"

func ErrorWhileAuthorizingRequest() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ErrorWhileAuthorizingRequestErr,
		Message: "error authorizing request",
	}
}

// The authorization server denied the request
const UnauthorizedErr Code = "UNAUTHORIZED"

func Unauthorized() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnauthorizedErr,
		Message: "authorization server denied request",
	}
}

// Unable to verify claim for delegation, claim can not be parsed
const UnableToVerifyClaimErr Code = "UNABLE_TO_VERIFY_CLAIM"

func UnableToVerifyClaim() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnableToVerifyClaimErr,
		Message: "unable to verify claim",
	}
}

// The requesting organization is not the organization found in the order
const RequestingOrganizationIsNotDelegateeErr Code = "REQUESTING_ORGANIZATION_IS_NOT_DELEGATEE"

func RequestingOrganizationIsNotDelegatee(msg string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    RequestingOrganizationIsNotDelegateeErr,
		Message: msg,
	}
}

// The delegator of this order does not have access to the service
const DelegatorDoesNotHaveAccessToServiceErr Code = "DELEGATOR_DOES_NOT_HAVE_ACCESS_TO_SERVICE"

func DelegatorDoesNotHaveAccessToService() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    DelegatorDoesNotHaveAccessToServiceErr,
		Message: "no access. delegator does not have access to the service for the public key in the claim",
	}
}

// Missing log record ID, the header 'X-NLX-Logrecord-Id' must be set with an unique ID for this request
const MissingLogRecordIDErr Code = "MISSING_LOG_RECORD_ID"

func MissingLogRecordID() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    MissingLogRecordIDErr,
		Message: "missing logrecord id",
	}
}

// General server error, see message for more information
const ServerErrorErr Code = "SERVER_ERROR"

func ServerError(errDetails error) *NLXNetworkError {
	message := "server error"

	if errDetails != nil {
		message = fmt.Sprintf("%s: %v", message, errDetails)
	}

	return &NLXNetworkError{
		Code:    ServerErrorErr,
		Message: message,
	}
}

// Outway is called with an invalid URL, see message for more information
const InvalidURLErr Code = "INVALID_URL"

func InvalidURL(msg string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    InvalidURLErr,
		Message: msg,
	}
}

// Proxy mode is disabled, enable it by setting the 'use-as-http-proxy' flag to resolve
const ProxyModeDisabledErr Code = "PROXY_MODE_DISABLED"

func ProxyModeDisabled(url string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ProxyModeDisabledErr,
		Message: fmt.Sprintf("please enable proxy mode by setting the 'use-as-http-proxy' flag to resolve: %s", url),
	}
}

// Outway got called with an invalid method, CONNECT method is not supported
const UnsupportedMethodErr Code = "UNSUPPORTED_METHOD"

func UnsupportedMethod() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnsupportedMethodErr,
		Message: "CONNECT method is not supported",
	}
}

// Unable to parse delegation metadata,
// check if 'X-NLX-Request-Delegator' and 'X-NLX-Request-Order-Reference' headers are set correctly.
const UnableToParseDelegationMetadataErr Code = "UNABLE_TO_PARSE_DELEGATION_METADATA"

func UnableToParseDelegationMetadata() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnableToParseDelegationMetadataErr,
		Message: "failed to parse delegation metadata",
	}
}

// Unable to setup management client to retrieve claim for order
const UnableToSetupManagementClientErr Code = "UNABLE_TO_SETUP_MANAGEMENT_CLIENT"

func UnableToSetupManagementClient() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnableToSetupManagementClientErr,
		Message: "unable to setup the external management client",
	}
}

// Order was not found
const OrderNotFoundErr Code = "ORDER_NOT_FOUND"

func OrderNotFound() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    OrderNotFoundErr,
		Message: "order not found",
	}
}

// The used order does not exist for your organization
const OrderDoesNotExistForYourOrganizationErr Code = "ORDER_DOES_NOT_EXIST_FOR_YOUR_ORGANIZATION"

func OrderDoesNotExistForYourOrganization() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    OrderDoesNotExistForYourOrganizationErr,
		Message: "order does not exist for your organization",
	}
}

// The order is revoked by the delegator
const OrderRevokedErr Code = "ORDER_REVOKED"

func OrderRevoked() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    OrderRevokedErr,
		Message: "order is revoked",
	}
}

// The order is expired
const OrderExpiredErr Code = "ORDER_EXPIRED"

func OrderExpired() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    OrderExpiredErr,
		Message: "the order has expired",
	}
}

// The order does not contain the service that was tried to access
const OrderDoesNotContainServiceErr Code = "ORDER_DOES_NOT_CONTAIN_SERVICE"

func OrderDoesNotContainService(serviceName string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    OrderDoesNotContainServiceErr,
		Message: fmt.Sprintf("order does not contain the service '%s'", serviceName),
	}
}

// Unable to request claim from delegator
const UnableToRequestClaimErr Code = "UNABLE_TO_REQUEST_CLAIM"

func UnableToRequestClaim(delegatorSerialNumber string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    UnableToRequestClaimErr,
		Message: fmt.Sprintf("unable to request claim from %s", delegatorSerialNumber),
	}
}

// Received an invalid claim from the delegator
const ReceivedInvalidClaimErr Code = "RECEIVED_INVALID_CLAIM"

func ReceivedInvalidClaim(delegatorSerialNumber string) *NLXNetworkError {
	return &NLXNetworkError{
		Code:    ReceivedInvalidClaimErr,
		Message: fmt.Sprintf("received an invalid claim from %s", delegatorSerialNumber),
	}
}

// Invalid data subject header, 'X-NLX-Request-Data-Subject' contains invalid data.
// Must be in 'key=value' format.
const InvalidDataSubjectHeaderErr Code = "INVALID_DATA_SUBJECT_HEADER"

func InvalidDataSubjectHeader() *NLXNetworkError {
	return &NLXNetworkError{
		Code:    InvalidDataSubjectHeaderErr,
		Message: "invalid data subject header",
	}
}

// Source
const (
	// The error originated from the Inway
	Inway Source = "inway"

	// The error originated from the Outway
	Outway Source = "outway"
)

// Location
// Check https://docs.nlx.io/support/common-errors for a graphical overview
const (
	// The error happened between the Inway and the API
	A1 Location = "A1"

	// The error happened between the Inway and the Authorization server
	IAS1 Location = "IAS1"

	// The error happened between the Client and the Outway
	C1 Location = "C1"

	// The error happened between the Directory Monitor and the Inway.
	// This means that the Inway couldn't be reached from the Directory Monitor and thus Outways cannot reach your Inway
	M1 Location = "M1"

	// The error happened between the Outway and Inway
	O1 Location = "O1"

	// The error happened between the Outway and the Authorization server
	OAS1 Location = "OAS1"
)

func (l Location) String() string {
	return string(l)
}

func (l Location) GoString() string {
	return l.String()
}

func (l Code) String() string {
	return string(l)
}

func (l Code) GoString() string {
	return l.String()
}

func (l Source) String() string {
	return string(l)
}

func (l Source) GoString() string {
	return l.String()
}
