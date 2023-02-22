#!/bin/bash

eval $(cat .env | sed 's/#.*//g')

migrate -database "postgres://$DB_USER:$DB_PASSWORD@localhost:$DB_PORT/$DB_NAME?sslmode=disable" -path sql up