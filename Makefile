SHELL=/bin/bash
ARCH=$(shell uname -r | sed 's/.*-\(.*\)/\1/')
OS=$(shell uname | tr '[:upper:]' '[:lower:]')

.PHONY: build test clean static-binary all-static-bin install

build:main.go
	go build -o viz .

test:
	go test -v ./internal/...

clean:
	rm -rf viz*
    rm -rf binary/viz*

# Use if you are using POSIX compliant shell
static-binary:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o viz-$(OS)-$(ARCH)

all-static-bin:
	GOOS=linux GOARCH=amd64 go build -o binary/viz-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o binary/viz-linux-arm64

	GOOS=darwin GOARCH=amd64 go build -o binary/viz-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o binary/viz-darwin-arm64

	GOOS=windows GOARCH=amd64 go build -o binary/viz-windows-amd64.exe
	GOOS=windows GOARCH=amd64 go build -o binary/viz-windows-amd64.exe

install:
	./scripts/install.sh
