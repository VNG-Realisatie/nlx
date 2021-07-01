ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

PKG_FILES=pkg/database/mock/database.go domain/directory/mock/repository.go

.PHONY: all
all: generate

.PHONY: generate
generate: ${PKG_FILES}

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f ${PKG_FILES}

.PHONY: dependencies
dependencies:
	go install github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

pkg/database/mock/database.go:
	@# Generate mock file
	mockgen -source pkg/database/database.go -package=mock -destination $@

domain/directory/mock/repository.go:
	@# Generate mock file
	mockgen -source domain/directory/repository.go -package=directory_mock -destination $@
