package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrServiceDoesNotExist = status.Error(codes.NotFound, "service does not exist")
)
