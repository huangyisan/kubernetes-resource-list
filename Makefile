.PHONY: build-linux build-osx build-windows clean

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

build-linux:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/kube-resource_linux

build-osx:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/kube-resource_darwin

build-windows:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/kube-resource_windows

clean:
	@if [ -f bin/kube-resource_linux ] ; then rm bin/kube-resource_linux ; fi
	@if [ -f bin/kube-resource_darwin ] ; then rm bin/kube-resource_darwin ; fi
	@if [ -f bin/kube-resource_windows ] ; then rm bin/kube-resource_windows ; fi