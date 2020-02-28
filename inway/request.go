// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inway

// RequestMetadata contains information about an incoming request.
// It is passed to ServiceEndpoint's handleRequest which uses it for implementation-specific proxying.
type RequestMetadata struct {
	requestPath            string
	requesterPublicKeyHash string
	requesterOrganization  string
}
