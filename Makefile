
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	GO111MODULE=on && go get -u github.com/go-kratos/kratos/tool/kratos

.PHONY: proto
proto:
	cd api;go generate

.PHONY: generate
generate:
	go get github.com/google/wire/cmd/wire
	cd internal/dao;go generate
	cd internal/service;go generate
	cd internal/service/collect;go generate
	cd internal/di;go generate

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fantastic-happiness cmd/*.go

.PHONY: check
check:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t fantastic-happiness:latest
