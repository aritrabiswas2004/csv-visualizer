SHELL=/bin/bash
ARCH=$(shell uname -r | sed 's/.*-\(.*\)/\1/')
OS=$(shell uname | tr '[:upper:]' '[:lower:]')

.PHONY: build test clean static-binary

build:main.go
	go build -o viz .

test:
	go test -v ./internal/...

clean:
	rm -rf viz*

static-binary:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o viz-$(OS)-$(ARCH)

