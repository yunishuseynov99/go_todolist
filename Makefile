include .env
export

export PROJECT_ROOT := $(CURDIR)

env-up:
	@docker compose --env-file .env up -d todoapp-postgres

env-down:
	# CHANGE IS HERE: Added --remove-orphans
	@docker compose --env-file .env down --remove-orphans

env-cleanup:
	@read -p "This will delete postgres data. Type Y to continue: " ans; \
	if [ "$$ans" = "Y" ]; then \
		# CHANGE IS HERE: Added --remove-orphans
		docker compose --env-file .env down --remove-orphans && \
		rm -rf out/pgdata && \
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

# REMOVED env-port-forward and env-port-close targets

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing variable action. Example: action=up"; \
		exit 1; \
	fi; \
	MSYS_NO_PATHCONV=1 docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go