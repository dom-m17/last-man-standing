.PHONY: gen gen-all start postgres createdb dropdb migrate-up migrate-down test tests

gen: ## Generate GraphQL code
	go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate --config gqlgen.yml

gen-all: ## Generate everything in the project
	go generate ./...

start:
	air 

postgres:
	docker run --name postgres17.4 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:17.4-alpine

createdb:
	docker exec -it postgres17.4 createdb --username=root --owner=root lms

dropdb:
	docker exec -it postgres17.4 dropdb --username=root --owner=root lms

migrate-up:
	migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose up

migrate-down:
	migrate -path sql/migration -database "postgresql://root:password@localhost:5432/lms?sslmode=disable" -verbose down

test: ## Run one particular test
	go test -v -cover -race -run=$(name)

tests: ## Run all tests in the project
	go test -v -cover -race ./...