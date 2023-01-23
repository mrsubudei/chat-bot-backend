include .env
export

.PHONY: up
up: ### Run docker-compose
	docker-compose up -d
