package inway

// RequestMetadata contains information about an incomming request.
// It is passed to ServiceEndpoint's handleRequest which uses it for implementation-specifc proxying.
type RequestMetadata struct {
	requestPath           string
	requesterOrganization string
}
