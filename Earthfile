VERSION 0.5

FROM registry.gitlab.com/commonground/nlx/earthly-base-image:latest

WORKDIR /src

all:
    BUILD +proto
    BUILD +mocks
    BUILD +sqlc
    BUILD +enums

proto:
    BUILD +proto-directory-api
    BUILD +proto-management-api
    BUILD +proto-management-api-client
    BUILD +proto-txlog-api
    BUILD +proto-inway-test
    BUILD +proto-outway
    BUILD +proto-grpc-errors

mocks:
    BUILD +mocks-management-api
    BUILD +mocks-txlog-api
    BUILD +mocks-common
    BUILD +mocks-directory-api

sqlc:
    BUILD +sqlc-txlog-api
    BUILD +sqlc-directory-api
    BUILD +sqlc-management-api

enums:
    BUILD +enums-permissions

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
            --openapiv2_opt json_names_for_fields=false \
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
            -I./external \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist \
            --openapiv2_opt json_names_for_fields=false \
            ./*.proto && \
        cd external && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist/external --go_opt=paths=source_relative \
            --go-grpc_out=/dist/external --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist/external --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist/external \
            --openapiv2_opt json_names_for_fields=false \
            ./*.proto

    RUN  npx @openapitools/openapi-generator-cli generate -i /dist/management.swagger.json -g openapi --additional-properties=outputFileName=management.swagger.json -o /openapi
    RUN  npx @openapitools/openapi-generator-cli generate -i /dist/external/external.swagger.json -g openapi --additional-properties=outputFileName=external.swagger.json -o /openapi

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/*.* AS LOCAL ./management-api/api/
    SAVE ARTIFACT /dist/external/*.* AS LOCAL ./management-api/api/external/
    SAVE ARTIFACT /openapi/management.swagger.json AS LOCAL ./management-api/api/
    SAVE ARTIFACT /openapi/external.swagger.json AS LOCAL ./management-api/api/external/

proto-management-api-client:
    FROM +proto-management-api
    COPY ./management-api/api /src/management-api/api
    COPY ./management-ui /src/management-ui

    WORKDIR ./management-ui
    RUN npx @openapitools/openapi-generator-cli generate

    SAVE ARTIFACT ./src/api AS LOCAL ./management-ui/src/api

proto-txlog-api:
    FROM +deps
    COPY ./txlog-api/api/*.proto /src/

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist \
            --openapiv2_opt json_names_for_fields=false \
            ./txlog.proto
    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/*.* AS LOCAL ./txlog-api/api/

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

proto-outway:
    FROM +deps
    COPY ./outway/api/*.proto /src/

    RUN mkdir -p /dist || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            ./outway.proto

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/*.* AS LOCAL ./outway/api/

proto-grpc-errors:
    FROM +deps
    COPY ./common/grpcerrors/errors/*.proto /src/

    RUN mkdir -p /dist/external || true && \
        protoc \
            -I. \
            -I/protobuf/include \
            -I/protobuf/googleapis \
            --go_out=/dist --go_opt=paths=source_relative \
            --go-grpc_out=/dist --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=/dist --grpc-gateway_opt=paths=source_relative \
            --openapiv2_out=/dist \
            --openapiv2_opt json_names_for_fields=false \
            ./*.proto

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/*.* AS LOCAL ./common/grpcerrors/errors/

mocks-management-api:
    FROM +deps
    COPY ./management-api /src/management-api
    COPY ./common /src/common
    COPY ./directory-api /src/directory-api
    COPY ./txlog-api /src/txlog-api
    COPY ./outway /src/outway

    RUN mkdir -p /dist || true
    WORKDIR /src/management-api

    RUN mockgen -source api/management_grpc.pb.go -destination /dist/management-api/api/mock/mock_management.go
    RUN mockgen -source api/external/external_grpc.pb.go -destination /dist/management-api/api/external/mock/mock_external.go
    RUN mockgen -source pkg/database/database.go -destination /dist/management-api/pkg/database/mock/mock_database.go
    RUN mockgen -destination /dist/management-api/pkg/directory/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/directory Client
    RUN mockgen -destination /dist/management-api/pkg/txlog/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/txlog Client
    RUN mockgen -destination /dist/management-api/pkg/outway/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/outway Client
    RUN mockgen -destination /dist/management-api/pkg/management/mock/mock_client.go go.nlx.io/nlx/management-api/pkg/management Client
    RUN mockgen -source pkg/auditlog/logger.go -destination /dist/management-api/pkg/auditlog/mock/mock_auditlog.go
    RUN mockgen -source pkg/txlogdb/database.go -destination /dist/management-api/pkg/txlogdb/mock/mock_database.go
    RUN mockgen -source pkg/oidc/authenticator.go -destination /dist/management-api/pkg/oidc/mock/mock_authenticator.go

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/management-api/api/mock/*.go AS LOCAL ./management-api/api/mock/
    SAVE ARTIFACT /dist/management-api/api/external/mock/*.go AS LOCAL ./management-api/api/external/mock/
    SAVE ARTIFACT /dist/management-api/pkg/database/mock/*.go AS LOCAL ./management-api/pkg/database/mock/
    SAVE ARTIFACT /dist/management-api/pkg/directory/mock/*.go AS LOCAL ./management-api/pkg/directory/mock/
    SAVE ARTIFACT /dist/management-api/pkg/txlog/mock/*.go AS LOCAL ./management-api/pkg/txlog/mock/
    SAVE ARTIFACT /dist/management-api/pkg/outway/mock/*.go AS LOCAL ./management-api/pkg/outway/mock/
    SAVE ARTIFACT /dist/management-api/pkg/management/mock/*.go AS LOCAL ./management-api/pkg/management/mock/
    SAVE ARTIFACT /dist/management-api/pkg/auditlog/mock/*.go AS LOCAL ./management-api/pkg/auditlog/mock/
    SAVE ARTIFACT /dist/management-api/pkg/txlogdb/mock/*.go AS LOCAL ./management-api/pkg/txlogdb/mock/
    SAVE ARTIFACT /dist/management-api/pkg/oidc/mock/*.go AS LOCAL ./management-api/pkg/oidc/mock/

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

mocks-txlog-api:
    FROM +deps
    COPY ./txlog-api /src/txlog-api

    RUN mkdir -p /dist || true
    WORKDIR /src/txlog-api

    RUN mockgen -source api/txlog_grpc.pb.go -package=mock -destination /dist/api/mock/mock_txlog.go
    RUN mockgen -source domain/record/repository.go -package=txlog_mock -destination /dist/domain/record/mock/repository.go

    RUN goimports -w -local "go.nlx.io" /dist/

    SAVE ARTIFACT /dist/api/mock/mock_txlog.go AS LOCAL ./txlog-api/api/mock/mock_txlog.go
    SAVE ARTIFACT /dist/domain/record/mock/repository.go AS LOCAL ./txlog-api/domain/record/mock/repository.go

sqlc-txlog-api:
    FROM +deps
    COPY ./txlog-api/adapters/storage/postgres/queries /src/txlog-api/adapters/storage/postgres/queries
    COPY ./txlog-api/adapters/storage/postgres/migrations/sql /src/txlog-api/adapters/storage/postgres/migrations/sql

    WORKDIR /src/txlog-api/adapters/storage/postgres/queries

    RUN /usr/bin/sqlc generate

    RUN goimports -w -local "go.nlx.io" /src/

    SAVE ARTIFACT /src/txlog-api/adapters/storage/postgres/queries/* AS LOCAL ./txlog-api/adapters/storage/postgres/queries/

sqlc-directory-api:
    FROM +deps
    COPY ./directory-api/adapters/storage/postgres/queries /src/directory-api/adapters/storage/postgres/queries
    COPY ./directory-api/migrations /src/directory-api/migrations

    WORKDIR /src/directory-api/adapters/storage/postgres/queries

    RUN /usr/bin/sqlc generate

    RUN goimports -w -local "go.nlx.io" /src/

    SAVE ARTIFACT /src/directory-api/adapters/storage/postgres/queries/* AS LOCAL ./directory-api/adapters/storage/postgres/queries/

sqlc-management-api:
    FROM +deps
    COPY ./management-api/adapters/storage/postgres/queries /src/management-api/adapters/storage/postgres/queries
    COPY ./management-api/migrations /src/management-api/migrations

    COPY ./management-api/pkg/oidc/pgsessionstore/queries /src/management-api/pkg/oidc/pgsessionstore/queries
    COPY ./management-api/pkg/oidc/pgsessionstore//migrations /src/management-api/pkg/oidc/pgsessionstore/migrations

    WORKDIR /src/management-api/adapters/storage/postgres/queries

    RUN /usr/bin/sqlc generate

    WORKDIR /src/management-api/pkg/oidc/pgsessionstore/queries

    RUN /usr/bin/sqlc generate

    RUN goimports -w -local "go.nlx.io" /src/

    SAVE ARTIFACT /src/management-api/adapters/storage/postgres/queries/* AS LOCAL ./management-api/adapters/storage/postgres/queries/
    SAVE ARTIFACT /src/management-api/pkg/oidc/pgsessionstore/queries/* AS LOCAL ./management-api/pkg/oidc/pgsessionstore/queries/

enums-permissions:
    FROM +deps

    COPY ./management-api/pkg/permissions/ /src/management-api/pkg/permissions/

    WORKDIR /src/

    RUN enumer -type=Permission -linecomment ./management-api/pkg/permissions

    SAVE ARTIFACT /src/management-api/pkg/permissions/*_enumer.go AS LOCAL ./management-api/pkg/permissions/
