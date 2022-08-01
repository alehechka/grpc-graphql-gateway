.PHONY: clean all

SPECDIR := "./spec"

all: build

build: init
	rm -rf ./spec/starwars
	for file in `find ./proto -name '*.proto'`; do \
		echo $$file; \
		protoc \
			-Iproto \
			-I../../proto/grpc-graphql-gateway \
			--plugin=../../dist/protoc-gen-graphql \
			--graphql_out=${SPECDIR} \
			--graphql_opt=paths=source_relative \
			--graphql_opt=field_json_name \
			--go_out=:${SPECDIR} \
			--go_opt=paths=source_relative \
			--go-grpc_out=:${SPECDIR}  \
			--go-grpc_opt=paths=source_relative \
			$$file; \
	done

init:
	if [ -L "./proto/google" ]; then rm ./proto/google; fi
	if [ -L "./proto/graphql" ]; then rm ./proto/graphql; fi
	cd ../../ && make

start: build
	cd app/graphql && GOOS=linux go build -o app main.go
	cd app/grpc && GOOS=linux go build -o app main.go
	docker-compose up -d --build

stop:
	docker-compose down

start_graphql:
	go run app/graphql/main.go

start_grpc:
	go run app/grpc/main.go

build_graphql:
	go build -o app/graphql/ app/graphql/main.go

build_grpc:
	go build -o app/grpc/ app/grpc/main.go

build_all: build build_grpc build_graphql