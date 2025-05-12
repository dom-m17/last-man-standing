.PHONY: gen gen-all

gen: ## Generate GraphQL code
	go get github.com/99designs/gqlgen; go run github.com/99designs/gqlgen generate --config gqlgen.yml

gen-all: ## Generate everything in the project
	go generate ./...