GIN_MODE=release
PG_URL=postgres://user:pass@localhost:5432/postgres

.PHONY: up
up: ### Run docker-compose
	docker-compose up --build

.PHONY: down
down: ### Run docker-compose
	docker-compose down