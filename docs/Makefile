
.PHONY: all
all: clean build

.PHONY: build
build:
	git submodule update --init themes/docdock
	hugo

.PHONY: clean
clean:
	rm -rf public
