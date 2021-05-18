ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

API_FILES=inspectionapi/mock/mock_directory_inspection_api.go
PKG_FILES=pkg/database/mock/mock_database.go

.PHONY: all
all: generate

.PHONY: generate
generate: ${API_FILES} ${PKG_FILES}

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f ${API_FILES} ${PKG_FILES}

.PHONY: dependencies
dependencies:
	go install github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

inspectionapi/mock/mock_directory_inspection_api.go:
	@# Generate mock file
	mockgen -source inspectionapi/inspectionapi_grpc.pb.go -package=mock -destination $@

pkg/database/mock/mock_database.go:
	@# Generate mock file
	mockgen -source pkg/database/database.go -package=mock -destination $@
