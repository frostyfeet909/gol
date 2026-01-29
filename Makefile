BACKEND_DIR := backend
MIGRATIONS_DIR := migrations

ifeq ($(OS),Windows_NT)
COPYCMD = powershell -NoProfile -Command Copy-Item -Force
else
COPYCMD = cp
endif


install:
	cd $(BACKEND_DIR) && make install
	cd $(MIGRATIONS_DIR) && make install
	$(COPYCMD) .env.example .env

db:
	docker-compose up --build -d
	cd $(MIGRATIONS_DIR) && make migrate-up

run:
	cd $(BACKEND_DIR) && make run