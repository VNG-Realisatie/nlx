#!/bin/sh
# Copyright © VNG Realisatie 2022
# Licensed under the EUPL


psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE nlx_management_org_a;
    CREATE DATABASE nlx_management_org_b;
    CREATE DATABASE nlx_management_org_c;
    CREATE DATABASE nlx_txlog_a;
    CREATE DATABASE nlx_txlog_b;
    CREATE DATABASE nlx_txlog_c;
EOSQL
