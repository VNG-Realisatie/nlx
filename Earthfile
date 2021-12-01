FROM registry.gitlab.com/commonground/nlx/earthly-base-image:latest

WORKDIR /src

all:
    BUILD +proto
    BUILD +mocks

proto:
    BUILD +proto-directory-api
    BUILD +proto-management-api
    BUILD +proto-inway-test

mocks:
    BUILD +mocks-management-api
    BUILD +mocks-common
    BUILD +mocks-directory-api

deps:
    COPY go.mod go.sum /src/

proto-directory-api:
    FROM +deps
    COPY ./directory-api/api/*.proto /src
    COPY ./scripts/fix-directoryapi_grpc.pb.go.sh /fix.sh

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist \
            --openapiv2_out=/dist \
            ./directoryapi.proto
    RUN goimports -w -local "go.nlx.io" /dist/

    RUN sh /fix.sh /dist/directoryapi_grpc.pb.go

    SAVE ARTIFACT /dist/* AS LOCAL ./directory-api/api/

proto-management-api:
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
    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/*.* AS LOCAL ./management-api/api/
    SAVE ARTIFACT /dist/external/*.* AS LOCAL ./management-api/api/external/

proto-inway-test:
    FROM +deps
    COPY ./inway/grpcproxy/test/*.proto /src

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            ./test.proto
    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/* AS LOCAL ./inway/grpcproxy/test/

mocks-management-api:
    FROM +deps
    COPY ./management-api /src/management-api
    COPY ./common /src/common
    COPY ./directory-api /src/directory-api

    RUN mkdir -p /dist || true
    WORKDIR /src/management-api

    RUN mockgen -source api/management_grpc.pb.go -destination /dist/management-api/api/mock/mock_management.go
    RUN mockgen -source api/external/external_grpc.pb.go -destination /dist/management-api/api/external/mock/mock_external.go
    RUN mockgen -source pkg/database/database.go -destination /dist/management-api/pkg/database/mock/mock_database.go
    RUN mockgen -destination /dist/management-api/pkg/directory/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/directory Client
    RUN mockgen -destination /dist/management-api/pkg/management/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/management Client
    RUN mockgen -source pkg/auditlog/logger.go -destination /dist/management-api/pkg/auditlog/mock/mock_auditlog.go
    RUN mockgen -source pkg/txlogdb/database.go -destination /dist/management-api/pkg/txlogdb/mock/mock_database.go

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/management-api/api/mock/*.go AS LOCAL ./management-api/api/mock/
    SAVE ARTIFACT /dist/management-api/api/external/mock/*.go AS LOCAL ./management-api/api/external/mock/
    SAVE ARTIFACT /dist/management-api/pkg/database/mock/*.go AS LOCAL ./management-api/pkg/database/mock/
    SAVE ARTIFACT /dist/management-api/pkg/directory/mock/*.go AS LOCAL ./management-api/pkg/directory/mock/
    SAVE ARTIFACT /dist/management-api/pkg/management/mock/*.go AS LOCAL ./management-api/pkg/management/mock/
    SAVE ARTIFACT /dist/management-api/pkg/auditlog/mock/*.go AS LOCAL ./management-api/pkg/auditlog/mock/
    SAVE ARTIFACT /dist/management-api/pkg/txlogdb/mock/*.go AS LOCAL ./management-api/pkg/txlogdb/mock/

mocks-common:
    FROM +deps
    COPY ./common /src/common

    RUN mkdir -p /dist || true
    WORKDIR /src/common

    RUN mockgen -source ./transactionlog/logger.go -destination /dist/mock_logger.go
    RUN goimports -w -local "go.nlx.io" /dist/
    SAVE ARTIFACT /dist/mock_logger.go AS LOCAL ./common/transactionlog/mock/mock_logger.go

mocks-directory-api:
    FROM +deps
    COPY ./directory-api /src/directory-api

    RUN mkdir -p /dist || true
    WORKDIR /src/directory-api

    RUN mockgen -source api/directoryapi_grpc.pb.go -package=mock -destination /dist/api/mock/mock_directory_api.go
    RUN mockgen -source domain/directory/storage/repository.go -package=directory_mock -destination /dist/domain/directory/storage/mock/repository.go

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/api/mock/mock_directory_api.go AS LOCAL ./directory-api/api/mock/mock_directory_api.go
    SAVE ARTIFACT /dist/domain/directory/storage/mock/repository.go AS LOCAL ./directory-api/domain/directory/storage/mock/repository.go