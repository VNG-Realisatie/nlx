FROM golang:1.16.4-alpine

WORKDIR /src

all:
    BUILD +proto
    BUILD +mocks

proto:
    BUILD +proto-directory-inspection-api
    BUILD +proto-directory-registration-api
    BUILD +proto-management-api
    BUILD +proto-inway-test

mocks:
    BUILD +mocks-management-api
    BUILD +mocks-common
    BUILD +mocks-directory-inspection-api
    BUILD +mocks-directory-registration-api

proto-deps:
    ENV PROTOBUF_VERSION=3.17.3

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

mocks-deps:
    COPY go.mod go.sum /src/

    RUN apk add --no-cache curl git unzip

    RUN go mod download

    RUN go install github.com/golang/mock/mockgen@v1.6.0

    SAVE IMAGE --cache-hint

proto-directory-inspection-api:
    FROM +proto-deps

    COPY ./directory-inspection-api/inspectionapi/*.proto /src

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist \
            --openapiv2_out=/dist --openapiv2_opt=json_names_for_fields=false \
            ./inspectionapi.proto

    RUN echo "package inspectionapi" > /dist/inspectionapi.swagger.json.go && \
        echo "const (" >> /dist/inspectionapi.swagger.json.go && \
        echo "SwaggerJSONDirectoryInspection = \`" >> /dist/inspectionapi.swagger.json.go && \
        cat /dist/inspectionapi.swagger.json >> /dist/inspectionapi.swagger.json.go && \
        echo "\`)" >> /dist/inspectionapi.swagger.json.go && \
        go fmt /dist/inspectionapi.swagger.json.go

    SAVE ARTIFACT /dist/* AS LOCAL ./directory-inspection-api/inspectionapi/

proto-directory-registration-api:
    FROM +proto-deps

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

proto-management-api:
    FROM +proto-deps

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

proto-inway-test:
    FROM +proto-deps

    COPY ./inway/grpcproxy/test/*.proto /src

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            ./test.proto

    SAVE ARTIFACT /dist/* AS LOCAL ./inway/grpcproxy/test/

mocks-management-api:
    FROM +mocks-deps

    COPY ./management-api /src/management-api
    COPY ./common /src/common
    COPY ./directory-inspection-api /src/directory-inspection-api
    COPY ./directory-registration-api /src/directory-registration-api

    RUN mkdir -p /dist || true
    WORKDIR /src/management-api

    RUN mockgen -source api/management_grpc.pb.go -destination /dist/management-api/api/mock/mock_management.go
    SAVE ARTIFACT /dist/management-api/api/mock/*.go AS LOCAL ./management-api/api/mock/

    RUN mockgen -source api/external/external_grpc.pb.go -destination /dist/management-api/api/external/mock/mock_external.go
    SAVE ARTIFACT /dist/management-api/api/external/mock/*.go AS LOCAL ./management-api/api/external/mock/

    RUN mockgen -source pkg/database/database.go -destination /dist/management-api/pkg/database/mock/mock_database.go
    SAVE ARTIFACT /dist/management-api/pkg/database/mock/*.go AS LOCAL ./management-api/pkg/database/mock/

    RUN mockgen -destination /dist/management-api/pkg/directory/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/directory Client
    SAVE ARTIFACT /dist/management-api/pkg/directory/mock/*.go AS LOCAL ./management-api/pkg/directory/mock/

    RUN mockgen -destination /dist/management-api/pkg/management/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/management Client
    SAVE ARTIFACT /dist/management-api/pkg/management/mock/*.go AS LOCAL ./management-api/pkg/management/mock/

    RUN mockgen -source pkg/auditlog/logger.go -destination /dist/management-api/pkg/auditlog/mock/mock_auditlog.go
    SAVE ARTIFACT /dist/management-api/pkg/auditlog/mock/*.go AS LOCAL ./management-api/pkg/auditlog/mock/

    RUN mockgen -source pkg/txlogdb/database.go -destination /dist/management-api/pkg/txlogdb/mock/mock_database.go
    SAVE ARTIFACT /dist/management-api/pkg/txlogdb/mock/*.go AS LOCAL ./management-api/pkg/txlogdb/mock/

mocks-common:
    FROM +mocks-deps

    COPY ./common /src/common

    RUN mkdir -p /dist || true
    WORKDIR /src/common

    RUN mockgen -source ./transactionlog/logger.go -destination /dist/mock_logger.go
    SAVE ARTIFACT /dist/mock_logger.go AS LOCAL ./common/transactionlog/mock/mock_logger.go

mocks-directory-inspection-api:
    FROM +mocks-deps

    COPY ./directory-inspection-api /src/directory-inspection-api

    RUN mkdir -p /dist || true
    WORKDIR /src/directory-inspection-api

    RUN mockgen -source inspectionapi/inspectionapi_grpc.pb.go -package=mock -destination /dist/inspectionapi/mock/mock_directory_inspection_api.go
    SAVE ARTIFACT /dist/inspectionapi/mock/mock_directory_inspection_api.go AS LOCAL ./directory-inspection-api/inspectionapi/mock/mock_directory_inspection_api.go

    RUN mockgen -source pkg/database/database.go -package=mock -destination /dist/pkg/database/mock/mock_database.go
    SAVE ARTIFACT /dist/pkg/database/mock/mock_database.go AS LOCAL ./directory-inspection-api/pkg/database/mock/mock_database.go

mocks-directory-registration-api:
    FROM +mocks-deps

    COPY ./directory-registration-api /src/directory-registration-api

    RUN mkdir -p /dist || true
    WORKDIR /src/directory-registration-api

    RUN mockgen -source domain/directory/repository.go -package=directory_mock -destination /dist/domain/directory/mock/repository.go
    SAVE ARTIFACT /dist/domain/directory/mock/repository.go AS LOCAL ./directory-registration-api/domain/directory/mock/repository.go

    RUN mockgen -source pkg/database/database.go -package=mock -destination /dist/pkg/database/mock/database.go
    SAVE ARTIFACT /dist/pkg/database/mock/database.go AS LOCAL ./directory-registration-api/pkg/database/mock/database.go
