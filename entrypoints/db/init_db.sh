#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --host "$POSTGRES_HOST" --port "$POSTGRES_PORT" --username "$POSTGRES_USER" --password "$POSTGRES_PASSWORD" --dbname "$POSTGRES_DB" <<-EOSQL
    DROP DATABASE IF EXISTS auth-db;
    CREATE DATABASE auth-db;
EOSQL
