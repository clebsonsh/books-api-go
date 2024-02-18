build:
	@go build -o bin/main main.go

run: build
	@./bin/main

migrate:
	@go run scripts/migrate.go

migrate.fresh: 
	@go run scripts/migrate.go fresh