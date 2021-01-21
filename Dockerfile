FROM golang:1.15

RUN GO111MODULE=off go get -v google.golang.org/protobuf/cmd/protoc-gen-go
RUN GO111MODULE=on GOBIN=/usr/local/bin go get \
  github.com/bufbuild/buf/cmd/buf \
  github.com/bufbuild/buf/cmd/protoc-gen-buf-check-breaking \
  github.com/bufbuild/buf/cmd/protoc-gen-buf-check-lint \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc \
  github.com/twitchtv/twirp/protoc-gen-twirp

RUN apt update && apt install -y unzip
RUN PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
 curl -LO $PB_REL/download/v3.13.0/protoc-3.13.0-linux-x86_64.zip && \
 unzip protoc-3.13.0-linux-x86_64.zip -d /protoc

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0 && \
    mv ./bin/golangci-lint /usr/local/bin


WORKDIR /app
ADD build.sh /
