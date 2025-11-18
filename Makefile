.PHONY: test migrate rollback status create

MIGRATIONS_DIR := ./internal/database/migrations
ENV_FILE := ./configs/.env
DATABASE_ENV_FILE := ./configs/database.env
IMAGE_NAME := cli-tasks-api
CONTAINER_NAME := minha-api

-include $(ENV_FILE) 
export $(shell sed 's/=.*//' $(ENV_FILE) 2>/dev/null)
-include $(DATABASE_ENV_FILE)
export $(shell sed 's/=.*//' $(DATABASE_ENV_FILE) 2>/dev/null)

container:
	docker run -d --network host --env-file $(DATABASE_ENV_FILE) --name $(CONTAINER_NAME) $(IMAGE_NAME)

image:
	docker build -t $(IMAGE_NAME) .

start-container:
	docker start $(CONTAINER_NAME) || true

stop-container:
	docker stop $(CONTAINER_NAME) || true

test:
	go test ./...

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