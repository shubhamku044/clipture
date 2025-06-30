.PHONY: help dev build up down logs clean test

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
dev: ## Start development environment
	docker-compose -f docker-compose.dev.yml up --build

dev-detached: ## Start development environment in detached mode
	docker-compose -f docker-compose.dev.yml up --build -d

dev-down: ## Stop development environment
	docker-compose -f docker-compose.dev.yml down

dev-logs: ## Show development logs
	docker-compose -f docker-compose.dev.yml logs -f

# Production
build: ## Build all services
	docker-compose build

up: ## Start production environment
	docker-compose up --build

up-detached: ## Start production environment in detached mode
	docker-compose up --build -d

down: ## Stop production environment
	docker-compose down

restart: ## Restart all services
	docker-compose restart

logs: ## Show production logs
	docker-compose logs -f

# Database
db-logs: ## Show database logs
	docker-compose logs -f postgres

db-shell: ## Connect to database shell
	docker-compose exec postgres psql -U postgres -d clipture

db-backup: ## Backup database
	docker-compose exec postgres pg_dump -U postgres clipture > backup_$$(date +%Y%m%d_%H%M%S).sql

# Backend
backend-shell: ## Connect to backend container shell
	docker-compose exec backend sh

backend-logs: ## Show backend logs
	docker-compose logs -f backend

backend-test: ## Run backend tests
	cd backend && go test ./...

# Frontend
frontend-shell: ## Connect to frontend container shell
	docker-compose exec app sh

frontend-logs: ## Show frontend logs
	docker-compose logs -f app

# Cleanup
clean: ## Clean up containers, networks, and volumes
	docker-compose down -v --remove-orphans
	docker system prune -f

clean-all: ## Clean up everything including images
	docker-compose down -v --remove-orphans --rmi all
	docker system prune -af

# Health checks
health: ## Check health of all services
	@echo "Checking service health..."
	@curl -f http://localhost:8080/health || echo "Backend: UNHEALTHY"
	@curl -f http://localhost:3000/health || echo "Frontend: UNHEALTHY"
	@curl -f http://localhost:3001/health || echo "Landing: UNHEALTHY"

# Setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	cp backend/.env.example backend/.env
	@echo "Please edit backend/.env with your configuration"
	@echo "Run 'make dev' to start development environment"

# Install dependencies
install: ## Install dependencies for all apps
	cd backend && go mod download
	cd apps/app && npm install
	cd apps/landing && npm install 