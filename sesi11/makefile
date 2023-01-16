
running-postgres:
	docker run -d --name goneo-db -p 5433:5432  -e POSTGRES_PASSWORD=goneo -e POSTGRES_USER=goneo -e POSTGRES_DB=goneo -v $PWD/db:/var/lib/postgresql/data postgres:alpine

build-image:
	docker build -f deploy/Dockerfile -t goneo:1.1-alpine .

run-image:
	docker run -d --name goneo -e PORT=7777 -e DB_HOST=goneo-db -e DB_PORT=5432 -e DB_USER=goneo -e DB_PASSWORD=goneo -e DB_DBNAME=goneo -p 7777:7777 goneo:1.1

create-network:
	docker network create goneo

connect-network:
	docker network connect goneo goneo
	docker network connect goneo goneo-db