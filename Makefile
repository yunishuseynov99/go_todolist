include .env
export

export PROJECT_ROOT := $(CURDIR)

env-up:
	@docker compose --env-file .env up -d todoapp-postgres

env-down:
	@docker compose --env-file .env down --remove-orphans

env-cleanup:
	@read -p "This will delete postgres data. Type Y to continue: " ans; \
	if [ "$$ans" = "Y" ]; then \
		docker compose --env-file .env down --remove-orphans && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "env cleaned"; \
	else \
		echo "cancelled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Missing variable seq. Example: seq=init"; \
		exit 1; \
	fi; \
	MSYS_NO_PATHCONV=1 docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing variable action. Example: action=up"; \
		exit 1; \
	fi; \
	set -a; . ./.env; set +a; \
	MSYS_NO_PATHCONV=1 docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://$${POSTGRES_USER}:$${POSTGRES_PASSWORD}@todoapp-postgres:5432/$${POSTGRES_DB}?sslmode=disable" \
		"$(action)"

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	set -a && . ./.env && set +a && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go