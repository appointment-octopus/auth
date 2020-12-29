#!/bin/bash

set -e

psql postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB

# psql -v ON_ERROR_STOP=1 --host "$POSTGRES_HOST" --port "$POSTGRES_PORT" --username "$POSTGRES_USER" --password "$POSTGRES_PASSWORD" --dbname "$POSTGRES_DB" <<-EOSQL
#     DROP DATABASE IF EXISTS auth-db;
#     CREATE DATABASE auth-db;
# EOSQL
