
.PHONY: all
all: vendor build

.PHONY: vendor
vendor:
	go get ./...

.PHONY: build
build:
	go build -o dist/bin/nlx-inway ./cmd/nlx-inway
