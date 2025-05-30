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

config:
	@docker compose -f docker-compose-test.yaml config

compose-up:
	@docker compose -f docker-compose-test.yaml up --build --exit-code-from app

compose-down:
	@docker compose -f docker-compose-test.yaml down

compose-logs:
	@docker compose -f docker-compose-test.yaml logs -f

compose-restart:
	@docker compose -f docker-compose-test.yaml up --build --force-recreate

# ----------------------------------------
# Tests
# ----------------------------------------

test: check-env compose-up compose-down

# ----------------------------------------
# Convenience
# ----------------------------------------

ps:
	docker compose ps
