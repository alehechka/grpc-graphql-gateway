FROM golang:1.18-alpine as go-builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

ARG VERSION="latest"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN cd protoc-gen-graphql ; go build -ldflags "-X main.version=${VERSION}"

FROM scratch

ARG VERSION="latest"

# Runtime dependencies
LABEL "build.buf.plugins.runtime_library_versions.0.name"="google.golang.org/protobuf"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="v1.28.0"
LABEL "build.buf.plugins.runtime_library_versions.1.name"="github.com/alehechka/grpc-graphql-gateway"
LABEL "build.buf.plugins.runtime_library_versions.1.version"="${VERSION}"

COPY --from=go-builder /app/protoc-gen-graphql/protoc-gen-graphql /

ENTRYPOINT [ "/protoc-gen-graphql" ]