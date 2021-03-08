#!/bin/bash

set -e

mockgen -source=./pkg/database/database.go -destination=./pkg/database/mock/database.go -package=mock
mockgen -source=./inspectionapi/inspectionapi_grpc.pb.go -destination=./inspectionapi/mock/directory-inspection-api.go -package=mock
