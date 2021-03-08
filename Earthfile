FROM golang:1.16-alpine

WORKDIR /src

proto:
    BUILD +directory-inspection-api
    BUILD +directory-registration-api
    BUILD +management-api

deps:
    ENV PROTOBUF_VERSION=3.15.5

    COPY go.mod go.sum /src/

    RUN apk add --no-cache protoc curl git unzip

    RUN go mod download

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

    SAVE IMAGE --cache-hint

directory-inspection-api:
    FROM +deps

    COPY ./directory-inspection-api/inspectionapi/*.proto /src

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist \
            --openapiv2_out=/dist \
            ./inspectionapi.proto

    RUN echo "package inspectionapi" > /dist/inspectionapi.swagger.json.go && \
        echo "const (" >> /dist/inspectionapi.swagger.json.go && \
        echo "SwaggerJSONDirectoryInspection = \`" >> /dist/inspectionapi.swagger.json.go && \
        cat /dist/inspectionapi.swagger.json >> /dist/inspectionapi.swagger.json.go && \
        echo "\`)" >> /dist/inspectionapi.swagger.json.go && \
        go fmt /dist/inspectionapi.swagger.json.go

    SAVE ARTIFACT /dist/* AS LOCAL ./directory-inspection-api/inspectionapi/

directory-registration-api:
    FROM +deps

    COPY ./directory-registration-api/registrationapi/*.proto /src

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist \
            --openapiv2_out=/dist \
            ./registrationapi.proto

    RUN echo "package registrationapi" > /dist/registrationapi.swagger.json.go && \
        echo "const (" >> /dist/registrationapi.swagger.json.go && \
        echo "SwaggerJSONDirectoryregistration = \`" >> /dist/registrationapi.swagger.json.go && \
        cat /dist/registrationapi.swagger.json >> /dist/registrationapi.swagger.json.go && \
        echo "\`)" >> /dist/registrationapi.swagger.json.go && \
        go fmt /dist/registrationapi.swagger.json.go

    SAVE ARTIFACT /dist/* AS LOCAL ./directory-registration-api/registrationapi/

management-api:
    FROM +deps

    COPY ./management-api/api/*.proto /src/
    COPY ./management-api/api/external/*.proto /src/external/

    RUN mkdir -p /dist/external || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist \
            ./management.proto && \
        cd external && \
        protoc \
            -I. \
            -I.. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist/external --go_opt=paths=source_relative \
            --go-grpc_out=/dist/external --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist/external --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist/external \
            ./external.proto

    SAVE ARTIFACT /dist/*.* AS LOCAL ./management-api/api/
    SAVE ARTIFACT /dist/external/*.* AS LOCAL ./management-api/api/external/
