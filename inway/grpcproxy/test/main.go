package test

import "google.golang.org/grpc"

func GetTestServiceDesc() *grpc.ServiceDesc {
	return &TestService_ServiceDesc
}
