postgres:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:15rc1-alpine
createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres15 createdb dropdb simple_bank
migrateup:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration --database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down
.PHONY: postgres createdb dropdb