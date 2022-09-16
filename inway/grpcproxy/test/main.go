// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package test

import "google.golang.org/grpc"

func GetTestServiceDesc() *grpc.ServiceDesc {
	return &TestService_ServiceDesc
}
