// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrServiceDoesNotExist = status.Error(codes.NotFound, "service does not exist")
)
