ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# ----------------------------------------
# Formatting & Linting
# ----------------------------------------

fmt:
	@echo "Formatting..."
	@gofmt -s -w ./src
	@goimports -w ./src

lint:
	@cd src && golangci-lint run

# ----------------------------------------
# Docker Compose Targets
# ----------------------------------------

check-env:
	@if [ -z "$(DB_USER)" ] || [ -z "$(DB_PASS)" ] || [ -z "$(DB_NAME)" ]; then \
		echo "❌ Missing DB_USER, DB_PASS, or DB_NAME!"; exit 1; \
	else \
		echo "✅ All DB env vars are set."; \
	fi

serve-config:
	@docker compose -f docker-compose.yaml config

test-setup:
	@docker compose -f docker-compose-test.yaml up --build --exit-code-from app

test-teardown:
	@docker compose -f docker-compose-test.yaml down

serve-setup:
	@docker compose -f docker-compose.yaml up --build

serve-teardown:
	@docker compose -f docker-compose.yaml down

# ----------------------------------------
# Run
# ----------------------------------------

test: check-env test-setup test-teardown
serve: check-env serve-setup

ps:
	docker compose ps
