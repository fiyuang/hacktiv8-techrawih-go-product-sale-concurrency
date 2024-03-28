.PHONY: clean build

dev:
	@echo "Running..."
	@go run main.go

dev-race-condition:
	@echo "Running..."
	@go run -race main.go

db-migrate-seed:
	@echo "Migrating & seeding datas..."
	@go run database/migrate/main.go
	@go run database/seeder/main.go

db-seed:
	@echo "Seeding datas..."
	@go run database/seeder/main.go