# Makefile
# шорткаты, экей сокращения команд. чтобы каждый раз в консоле не писать длинные команды. 
# миграции - команды для работы с базой, подключение к ней. Контроль версий за обновлений в базе данных.  

DB_URL = host=localhost user=postgres password=root dbname=FilmsCatalog port=5432 sslmode=disable
MIGRATIONS_DIR = migrations
DRIVER = postgres


# --- Миграции ---
migrate-up:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_URL)" status

create-migration:
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql

up: migrate-up
down: migrate-down
status: migrate-status
new: create-migration
