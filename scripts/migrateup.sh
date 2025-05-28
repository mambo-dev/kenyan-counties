#!/bin/bash

if [ -f .env ]; then
    source .env
fi

# Append authToken to the DATABASE_URL
FULL_DATABASE_URL="${DATABASE_URL}?authToken=${TAUTH_TOKEN}"

cd sql/schema
goose turso "$FULL_DATABASE_URL" up
