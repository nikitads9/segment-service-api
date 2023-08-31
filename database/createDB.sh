#!/bin/bash

export DB_HOST="postgres"
export DB_PORT="5432"
export DB_NAME="segments_db"
export DB_USER="postgres"
export DB_PASSWORD="avito_pass"
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

pgsql "${PG_DSN}" 