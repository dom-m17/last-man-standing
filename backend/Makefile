postgres:
	docker run --name postgres17.4 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.4-alpine

createdb:
	docker exec -it postgres17.4 createdb --username=root --owner=root lms

dropdb:
	docker exec -it postgres17.4 dropdb --username=root --owner=root lms

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown