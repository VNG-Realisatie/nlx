ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

.PHONY: all
all: migrations

.PHONY: migrations
migrations: db/migrations.go

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f db/migrations.go

.PHONY: dependencies
dependencies:
	go install github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

db/migrations.go:
	@# Generate migrations Go file
	go-bindata -prefix db/migrations -pkg db -o db/migrations.go db/migrations/...
