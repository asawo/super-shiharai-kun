# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

.PHONY: build
build:
	go build -o ./bin/server ./cmd/server 

.PHONY: run
run: build
	./bin/server

# ==================================================================================== #
# TESTS
# ==================================================================================== #
.PHONY: test
test:
	go test -v -race ./service/... ./http/...

# ==================================================================================== #
# DEBUG
# ==================================================================================== #

# Default values for local debugging
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= invoices-user
DB_NAME ?= invoices
DB_PASSWORD ?= p4ssw0rD

.PHONY: dump-db
dump-db:
	@echo "Dumping current database contents into db/seed/seed.sql..."
	@PGPASSWORD=$(DB_PASSWORD) pg_dump -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) > db/seed/seed.sql
	@echo "Database dump completed."

.PHONY: seed-db
seed-db:
	@echo "Restoring database from seed.sql..."
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -f db/seed/seed.sql
	@echo "Database restore completed."

.PHONY: drop-all-tables
drop-all-tables:
	@echo "Clearing all data from $(TABLE_NAME) table..."
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
	@echo "Data cleared from $(TABLE_NAME) table."
