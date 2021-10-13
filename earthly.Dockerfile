FROM golang:1.17.1-alpine

WORKDIR /src

 # General dependencies
RUN apk add --no-cache curl git unzip

# Proto dependencies
ENV PROTOBUF_VERSION=3.17.3

RUN apk add --no-cache protoc curl git unzip

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip && \
    unzip -q -d /protobuf protoc-${PROTOBUF_VERSION}-linux-x86_64.zip 'include/*' && \
    rm protoc-${PROTOBUF_VERSION}-linux-x86_64.zip

RUN curl -LO https://github.com/googleapis/googleapis/archive/master.zip && \
    unzip -q master.zip 'googleapis-master/google/api/*' && \
    mv googleapis-master /protobuf/googleapis && \
    rm master.zip

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.3.0 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.3.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

# Mock dependencies
RUN go install github.com/golang/mock/mockgen@v1.6.0
