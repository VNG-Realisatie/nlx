#
# Copyright © VNG Realisatie 2021
# Licensed under the EUPL
#

#!/bin/bash

export TEST_POSTGRES_DSN=postgres://localhost:5433/postgres?sslmode=disable

# empty database before running tests
psql -Atx $TEST_POSTGRES_DSN -c "drop schema if exists directory cascade;drop table schema_migrations;"

go test ./...
