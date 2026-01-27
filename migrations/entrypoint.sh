#! /usr/bin/env sh
migrate -path ./migrations -database "$DB_URL" up