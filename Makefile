.PHONY: build-linux build-osx build-windows clean

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

build-linux:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/linux/kube-resource

build-osx:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/darwin/kube-resource

build-windows:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/win/kube-resource

clean:
	@if [ -f bin/linux/kube-resource ] ; then rm bin/linux/ ; fi
	@if [ -f bin/darwin/kube-resource ] ; then rm bin/darwin/ ; fi
	@if [ -f bin/win/kube-resource ] ; then rm bin/win/ ; fi