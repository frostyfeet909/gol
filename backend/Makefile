ifeq ($(OS),Windows_NT)
COPYCMD = powershell -NoProfile -Command Copy-Item -Force
else
COPYCMD = cp
endif

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

install:
	go mod download
	go install github.com/swaggo/swag/cmd/swag@v1.16.6
	go install github.com/air-verse/air@v1.64.4
	$(COPYCMD) ../.env.example .env

run:
	swag init -g .\cmd\api\swagger.go -o .\docs
	go run -tags swagger,dotenv ./cmd/api
