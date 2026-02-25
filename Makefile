.PHONY: build run down migrate-up migrate-down migrate-create

-include .env

CUR_MIGRATION_DIR=$(MIGRATION_DIR)
MIGRATION_DSN="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL_MODE)"

# Собрать все сервисы без кеша
build:
	docker-compose build --no-cache

# Запустить все сервисы в фоновом режиме
up:
	docker-compose up -d

# Остановить все сервисы
down:
	docker-compose down

# Применить миграции
migrate-up:
	@migrate -database $(MIGRATION_DSN) -path $(CUR_MIGRATION_DIR) up

# Откатить последнюю миграцию
migrate-down:
	@migrate -database $(MIGRATION_DSN) -path $(CUR_MIGRATION_DIR) down -all

# Создать новую миграцию (использование: make migrate-create NAME=имя_миграции)
migrate-create:
	@[ -n "$(NAME)" ] || (echo "NAME=имя_миграции" && exit 1)
	migrate create -ext sql -dir $(CUR_MIGRATION_DIR) -seq $(NAME)

# Замер: 5 запросов статистики БЕЗ кеша (каждый раз тяжёлый SQL + pg_sleep 0.1s)
benchmark-stats-no-cache:
	@echo "=== GET /api/v1/users/1/stats (без кеша) — 5 запросов ==="
	@for i in 1 2 3 4 5; do curl -w "  запрос $$i: %{time_total}s\n" -o /dev/null -s "http://localhost:8080/api/v1/users/1/stats"; done

# Замер: прогрев кеша + 5 запросов статистики С кешем
benchmark-stats-cached:
	@echo "=== Прогрев кеша (1 запрос) ==="
	@curl -w "  время прогрева: %{time_total}s\n" -o /dev/null -s "http://localhost:8080/api/v1/users/1/stats/cached"
	@echo "=== GET /api/v1/users/1/stats/cached — 5 запросов из Redis ==="
	@for i in 1 2 3 4 5; do curl -w "  запрос $$i: %{time_total}s\n" -o /dev/null -s "http://localhost:8080/api/v1/users/1/stats/cached"; done

# Сравнение: вывести оба замера подряд
benchmark: benchmark-stats-no-cache benchmark-stats-cached
	@echo ""
	@echo "Итог: запросы из кеша должны быть на порядок быстрее (десятки ms vs сотни ms)."