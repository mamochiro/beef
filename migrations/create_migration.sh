#!/bin/bash
MIGRATION_NAME=$1

docker run \
  -v $PWD/sql:/migrations \
  --network host \
  migrate/migrate \
  -path=/migrations \
  create --dir /migrations -ext sql $MIGRATION_NAME
