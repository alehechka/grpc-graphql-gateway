FROM golang:1.18-alpine as go-builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

ARG RELEASE_VERSION

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN cd protoc-gen-graphql ; go build -ldflags "-X main.version=${RELEASE_VERSION}"

FROM scratch

ARG RELEASE_VERSION
ARG GO_PROTOBUF_RELEASE_VERSION
ARG GO_GRPC_RELEASE_VERSION

# Runtime dependencies
LABEL "build.buf.plugins.runtime_library_versions.0.name"="github.com/alehechka/grpc-graphql-gateway"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="${RELEASE_VERSION}"
LABEL "build.buf.plugins.runtime_library_versions.1.name"="google.golang.org/protobuf"
LABEL "build.buf.plugins.runtime_library_versions.1.version"="${GO_PROTOBUF_RELEASE_VERSION}"
LABEL "build.buf.plugins.runtime_library_versions.2.name"="google.golang.org/grpc"
LABEL "build.buf.plugins.runtime_library_versions.2.version"="${GO_GRPC_RELEASE_VERSION}"

COPY --from=go-builder /app/protoc-gen-graphql/protoc-gen-graphql /

ENTRYPOINT [ "/protoc-gen-graphql" ]