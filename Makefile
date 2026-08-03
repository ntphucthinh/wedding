.PHONY: help backend-up backend-down backend-restart backend-ps backend-build backend-build-nc backend-gen backend-codegen backend-migrate-up backend-migrate-down backend-migrate-refresh backend-migrate-fresh backend-seeder backend-shell frontend-help

help:
	@echo "Available commands:"
	@echo "  make backend-up           Start backend services"
	@echo "  make backend-down         Stop backend services"
	@echo "  make backend-build        Build backend services"
	@echo "  make backend-gen          Generate backend code"
	@echo "  make backend-shell        Open shell in the API container"
	@echo "  make frontend-help        Show frontend folder note"

backend-up:
	@make -C backend up

backend-down:
	@make -C backend down

backend-restart:
	@make -C backend restart

backend-ps:
	@make -C backend ps

backend-build:
	@make -C backend build

backend-build-nc:
	@make -C backend build-nc

backend-gen:
	@make -C backend gen

backend-codegen:
	@make -C backend codegen

backend-migrate-up:
	@make -C backend migrate-up

backend-migrate-down:
	@make -C backend migrate-down

backend-migrate-refresh:
	@make -C backend migrate-refresh

backend-migrate-fresh:
	@make -C backend migrate-fresh

backend-seeder:
	@make -C backend seeder

backend-shell:
	docker compose -f backend/docker-compose.yml exec wedding-api sh

frontend-help:
	@echo "Frontend folder is ready for a new app. Add your React/Vue/Next.js project here."
