GO_VERSION="1.16"
ARCH ?= $(shell go env GOOS)-$(shell go env GOARCH)


clean:
	rm -rf ~/dev/tickers/server/bin/*

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ~/dev/tickers/bin/tickers-amd64-linux
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ~/dev/tickers/bin/tickers-amd64.darwin
build-win:
	GOOS=windows GOARCH=amd64 go build -o ~/dev/tickers/bin/tickers-amd64.exe

all: clean build-linux build-mac build-win