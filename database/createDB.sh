#!/bin/bash

export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="segments_db"
export DB_USER="postgres"
export DB_PASSWORD="avito_pass"
export DB_SSL=disable

psql -h ${DB_HOST} -p ${DB_PORT} -U ${DB_USER} ${DB_NAME} < client-segment-db.sql