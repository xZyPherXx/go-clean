DB_CONTAINER=postgres
DB_USER=pwaadmin
DB_NAME=pwadb

up: 
	@docker compose up -d

down: 
	@docker compose down

build: 
	@docker compose up -d --build

restart: 
	@docker compose down && docker compose up -d

exe: 
	@docker exec -it $(IMG) /bin/sh

psql: 
	@docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

get:
	@curl -X GET localhost:8000/api/products && echo

getId:
	@curl -X GET localhost:8000/api/products/$(ID) && echo

delId:
	@curl -X DELETE localhost:8000/api/products/$(ID) && echo
