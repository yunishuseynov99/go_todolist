include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	docker compose up todoapp-postgres

env-down:
	docker compose down todoapp-postgres