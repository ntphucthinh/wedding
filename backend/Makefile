.PHONY: gen db-reset migrate-up run tools

build:
	docker compose build

build-nc:
	docker compose build --no-cache

up:
	@make set-up-githooks
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

ps:
	docker compose ps

set-up-githooks:
	mkdir -p .githooks
	git config --local core.hooksPath .githooks
	chmod -R +x .githooks/

# Generate backend + frontend (Alias for codegen)
gen:
	@make codegen

# Install requirement tools
tools:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/air-verse/air@v1.61.1
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

db-reset:
	docker compose exec -T wedding-db mysql -uroot -proot -e "DROP DATABASE IF EXISTS wedding_db; CREATE DATABASE wedding_db;"

migrate-up:
	migrate -path src/db/migrations -database "mysql://root:root@tcp(localhost:3309)/wedding_db" up

# run:
# 	air -c .air.toml

air:
	docker compose exec -e GOTOOLCHAIN=auto wedding-api air -c .air.toml

lint:
	docker compose exec wedding-api golangci-lint run

wire:
	docker compose exec -e GOTOOLCHAIN=auto wedding-api wire gen ./internal/di/wire.go
	docker compose exec -e GOTOOLCHAIN=auto wedding-api wire gen -output_file_prefix=test_ ./internal/di/wire-test.go

# OpenAPI Bundle
bundle:
	mkdir -p src/generated/api src/generated/openapi/openapi
	docker compose exec wedding-node openapi-generator-cli generate -c wedding-openapitools.json

codegen:
	@make bundle
	docker compose exec wedding-api oapi-codegen -config config/server.yaml generated/openapi/openapi/openapi.yaml
	docker compose exec wedding-api oapi-codegen -config config/models.yaml generated/openapi/openapi/openapi.yaml
# 	@make codegen-ts

codegen-ts:
	docker compose exec wedding-node npx openapi2aspida -i=./generated/openapi/openapi/openapi.yaml -o=./generated/backend-api
	@if [ -d "../frontend" ]; then \
		echo "Moving backend-api files to ../frontend..."; \
		cp -r src/generated/backend-api/@types/ ../frontend/src/generated/backend-api/@types/ 2>/dev/null || true; \
	fi
	docker compose exec wedding-node rm -rf generated/backend-api

#--------------------- migration start
migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: name parameter is required. Usage: make migration name=<migration_name>"; \
		exit 1; \
	fi
	docker compose exec wedding-api migrate create -ext sql -dir db/migrations $(name)

migrate-up:
	docker compose exec wedding-api migrate -path db/migrations -database "mysql://root:root@tcp(wedding-db)/wedding_db" up

migrate-down:
	docker compose exec wedding-api migrate -path db/migrations -database "mysql://root:root@tcp(wedding-db)/wedding_db" down -all

migrate-refresh:
	@make migrate-down
	@make migrate-up

migrate-fresh:
	docker compose exec wedding-db mysql --user=root --password=root -e 'DROP DATABASE IF EXISTS `wedding_db`;'
	docker compose exec wedding-db mysql --user=root --password=root -e 'CREATE DATABASE `wedding_db`;'
	@make migrate-up

seeder:
	docker compose exec -e GOTOOLCHAIN=auto wedding-api go run db/seed/seed.go