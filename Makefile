.PHONY: up
up: 
	docker-compose up --build 
	
.PHONY: down
down: 
	docker image prune --filter label=stage=build -f
	docker-compose down
