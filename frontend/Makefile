.PHONY: build build-nc up down restart ps shell install dev build-prod

build:
	docker compose build

build-nc:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

ps:
	docker compose ps

shell:
	docker compose exec wedding-frontend ash

install:
	docker compose run --rm wedding-frontend npm install

dev:
	docker compose run --rm -p 1301:1301 wedding-frontend npm run dev

build-prod:
	docker compose run --rm wedding-frontend npm run build
