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
	$(COPYCMD) ../.env.example .env

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path ./migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)
