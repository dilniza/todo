# Variables
DB_URI=mongodb://localhost:27017/todoapp

# Migrations
migration-up:
	go run cmd/migrate.go

# Start server
run:
	go run cmd/main.go

.PHONY: migration-up run
