ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

API_FILES=api/mock/mock_management.go api/external/mock/mock_external.go
PKG_FILES=pkg/database/mock/mock_database.go pkg/directory/mock/mock_client.go pkg/management/mock/mock_client.go pkg/auditlog/mock/mock_auditlog.go pkg/txlogdb/mock/mock_database.go

.PHONY: all
all: generate migrations

.PHONY: migrations
migrations: db/migrations.go

.PHONY: generate
generate: ${API_FILES} ${PKG_FILES}

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f ${API_FILES} ${PKG_FILES} db/migrations.go

.PHONY: dependencies
dependencies:
	GO111MODULE=off go get \
		github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

api/mock/mock_management.go:
	@# Generate mock file
	mockgen -source api/management_grpc.pb.go -destination $@

api/external/mock/mock_external.go:
	@# Generate mock file
	mockgen -source api/external/external_grpc.pb.go -destination $@

pkg/database/mock/mock_database.go:
	@# Generate mock file
	mockgen -source pkg/database/database.go -destination $@

pkg/directory/mock/mock_client.go:
	@# Generate mock file
	mockgen -destination pkg/directory/client.go go.nlx.io/nlx/management-api/pkg/directory Client

pkg/management/mock/mock_client.go:
	@# Generate mock file
	mockgen -destination pkg/management/client.go go.nlx.io/nlx/management-api/pkg/management Client

pkg/auditlog/mock/mock_auditlog.go:
	@# Generate mock file
	mockgen -source pkg/auditlog/logger.go -destination $@

pkg/txlogdb/mock/mock_database.go:
	@# Generate mock file
	mockgen -source pkg/txlogdb/database.go -destination $@

db/migrations.go:
	@# Generate migrations Go file
	go-bindata -prefix db/migrations -pkg db -o db/migrations.go db/migrations/...
