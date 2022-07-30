.PHONY: build

build:
	rm -rf greeter/*
	protoc \
	  -I. \
		-I../../proto/grpc-graphql-gateway \
		--plugin=../../dist/protoc-gen-graphql \
	  	--go_out=:./greeter \
		--go-grpc_out=./greeter \
		--graphql_out=verbose:./greeter \
	  greeter.proto

start_grpc:
	go run service/main.go

start_graphql:
	go run gateway/main.go

build_grpc:
	go build -o service/ service/main.go 

build_graphql:
	go build -o gateway/ gateway/main.go 

build_all: build build_grpc build_graphql