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

dcub:
	docker-compose up --build -d