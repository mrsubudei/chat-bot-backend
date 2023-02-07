include .env
export

.PHONY: up
up: ### Run docker-compose
	docker-compose up

.PHONY: cover
cover: 
	   go test ./... -coverprofile=coverage.out
	   go tool cover -html=coverage.out
