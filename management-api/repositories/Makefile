SOURCES = account.go

ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif


MOCKGEN = $(GOPATH)/bin/mockgen
MOCKGEN_URI = github.com/golang/mock/mockgen

MOCKS = $(patsubst %,mock/mock_%,$(SOURCES))

.PHONY: all
all: $(MOCKS)

$(MOCKGEN):
	go install $(MOCKGEN_URI)

mock/mock_%.go: %.go $(MOCKGEN)
	@# Generate mocks
	$(MOCKGEN) -source $<  -destination $@

.PHONY: clean
clean:
	@# Remove generated files
	rm -vf mock/mock_*.go

.PHONY: clean-dependencies
clean-dependencies:
	go clean -i $(MOCKGEN_URI)
