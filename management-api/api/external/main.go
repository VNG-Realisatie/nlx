// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package external

import "google.golang.org/grpc"

// GetAccessRequestServiceDesc returns the service specification for the AccessRequestService service
func GetAccessRequestServiceDesc() *grpc.ServiceDesc {
	return &AccessRequestService_ServiceDesc
}

// GetDelegationServiceDesc returns the service specification for the DelegationService service
func GetDelegationServiceDesc() *grpc.ServiceDesc {
	return &DelegationService_ServiceDesc
}
