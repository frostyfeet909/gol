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
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.19.1
	go install github.com/swaggo/swag/cmd/swag@v1.16.6
	$(COPYCMD) ../.env.example .env

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path ./migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

run:
	swag init -g .\cmd\api\swagger.go -o .\docs
	go run -tags swagger,dotenv ./cmd/api
