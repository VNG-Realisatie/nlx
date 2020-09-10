package test

import grpc "google.golang.org/grpc"

func GetTestServiceDesc() *grpc.ServiceDesc {
	return &_TestService_serviceDesc
}
