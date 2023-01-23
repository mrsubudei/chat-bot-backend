include .env
export

.PHONY: up
up: ### Run docker-compose
	docker-compose up --build -d postgres && docker-compose logs -f
