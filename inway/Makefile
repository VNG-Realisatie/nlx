
.PHONY: all
all: clean vendor build

.PHONY: clean
clean:
	rm -rf dist

.PHONY: vendor
vendor:
	go get ./...

.PHONY: build
build:
	go build -o dist/bin/nlx-inway ./cmd/nlx-inway
