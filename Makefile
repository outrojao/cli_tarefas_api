.PHONY: migrate rollback status create

MIGRATIONS_DIR := ./internal/database/migrations
ENV_FILE := ./configs/.env

-include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE) 2>/dev/null)

migrate:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

rollback:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

create:
	@if [ -z "$(name)" ]; then \
		echo "Uso: make create name=nome_da_migration"; \
	else \
		goose -dir $(MIGRATIONS_DIR) create $(name) sql; \
	fi