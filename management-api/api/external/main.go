// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package external

import "google.golang.org/grpc"

// GetAccessRequestServiceDesc returns the service specification for the AccessRequestService service
func GetAccessRequestServiceDesc() *grpc.ServiceDesc {
	return &_AccessRequestService_serviceDesc
}
