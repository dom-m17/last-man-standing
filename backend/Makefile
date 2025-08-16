.PHONY: sqlc gen gen-all start postgres createdb dropdb migrate-up migrate-down test tests seeds reset-db

sqlc:
	sqlc generate

gen: ## Generate GraphQL code
	go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate --config gqlgen.yml

gen-all: ## Generate everything in the project
	go generate ./...

run:
	air 

postgres:
	docker run --name lms-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.4-alpine

createdb:
	docker exec -it lms-postgres createdb --username=root --owner=root lms

dropdb:
	docker exec -it lms-postgres dropdb --username=root --owner=root lms

migrate-up:
	migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose up

migrate-down:
	migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose down

reset-db:
	migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose down && migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose up

test: ## Run one particular test
	docker-compose up -d; go test -v -cover -race -run=$(name)

tests: ## Run all tests in the project
	docker-compose up -d; go test -v -cover -race ./...

# TODO: improve this so that it can run multiple files?
seeds: ## Run the seed data
	docker exec -i lms-postgres psql -U root -d lms < ./sql/seeds/seed.sql