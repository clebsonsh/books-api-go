build:
	@go build -o bin/main main.go

run: build
	@./bin/main

migrate.fresh.seed: build
	@./bin/main migrate:fresh:seed

migrate.fresh: build
	@./bin/main migrate:fresh

migrate: build
	@./bin/main migrate

seed: build
	@./bin/main seed
