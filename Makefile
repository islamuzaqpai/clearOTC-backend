include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d clearotc-postgres

env-down:
	@docker compose down clearotc-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down clearotc-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения удалены"; \
	else \
		echo "Очистка отменена"; \
	fi


env-port-forward:
	@docker compose up -d port-forward

env-port-close:
	@docker compose down port-forward

migrate-create:
	@if [ -z "$(seq)" ]; then \
	   echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
	   exit 1; \
	fi; \
	MSYS_NO_PATHCONV=1 docker compose run --rm clearotc-postgres-migrate \
	   create \
	   -ext sql \
	   -dir /migrations \
	   -seq "$(seq)"

migrate-up:
	@make migration-action action=up

migrate-down:
	@make migration-action action=down

migration-action:
	@MSYS_NO_PATHCONV=1 docker compose run --rm clearotc-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@clearotc-postgres:5432/${POSTGRES_DB}?sslmode=disable&search_path=public" \
		"$(action)"

clearotc-run:
	go run cmd/main.go