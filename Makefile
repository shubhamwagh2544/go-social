# Makefile for database migrations using golang-migrate

MIGRATE_BIN ?= migrate
DB_URL ?= "postgres://postgres:postgres@localhost/go_social?sslmode=disable"
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migrate-create:
	$(MIGRATE_BIN) create -ext sql -dir $(MIGRATIONS_PATH) -seq $$name

.PHONY: migrate-up
migrate-up:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database $(DB_URL) up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE_BIN) -path $(MIGRATIONS_PATH) -database $(DB_URL) down 1
