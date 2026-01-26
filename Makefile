.PHONY: build test clean

build:main.go
	go build -o viz .

test:
	go test -v ./internal/...

clean:
	rm -rf viz
