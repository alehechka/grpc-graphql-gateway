.PHONY: command clean

GRAPHQL_CMD=protoc-gen-graphql
VERSION=$(or ${tag}, dev)
UNAME:=$(shell uname)
GOPATH:=$(shell go env GOPATH)

ifeq ($(UNAME), Darwin)
	PROTOPATH := $(shell brew --prefix protobuf)/include
endif
ifeq ($(UNAME), Linux)
	PROTOPATH := /usr/local/include
endif

command: plugin clean
	cd ${GRAPHQL_CMD} && \
		go build \
			-ldflags "-X main.version=${VERSION}" \
			-o ../dist/${GRAPHQL_CMD}

plugin:
	cd proto ; rm -rf gen ; buf generate

lint:
	golangci-lint run

test:
	cd runtime/testpb ; make build
	go test ./... --cover

build: test plugin

clean:
	rm -rf ./dist/*

all: clean build
	cd ${GRAPHQL_CMD} && GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o ../dist/${GRAPHQL_CMD}.darwin
	cd ${GRAPHQL_CMD} && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o ../dist/${GRAPHQL_CMD}.linux

compile:
	cd ${GRAPHQL_CMD} && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o ${GOPATH}/bin/${GRAPHQL_CMD}

compile_mac:
	cd ${GRAPHQL_CMD} && GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.version=${VERSION}" -o ${GOPATH}/bin/${GRAPHQL_CMD}

build_examples:
	cd example/starwars ; make build_all
	cd example/greeter ; make build_all