# Define variables
PORT := $(shell grep PORT .env | cut -d '=' -f2)
MAIN := cmd/myapp/main.go
LOG := server.log
PIDFILE := .pidfile

# Target to run the server in detached mode
run:
	@echo "Starting server on port $(PORT) in detached mode..."
	@nohup go run $(MAIN) > $(LOG) 2>&1 & echo $$! > $(PIDFILE)
	@echo "Server started. Check $(LOG) for output."

# Target to stop the server
stop:
	@if [ -f $(PIDFILE) ]; then \
		PID=$$(cat $(PIDFILE)); \
		kill $$PID && rm $(PIDFILE); \
		echo "Server stopped."; \
	else \
		echo "No server is running."; \
	fi

# Target to view the server log
log:
	@tail -f $(LOG)

# Target to install dependencies
install:
	@go get github.com/joho/godotenv
	@go get github.com/gorilla/mux
	@go get github.com/mattn/go-sqlite3
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@echo "Dependencies installed."

# Target to migrate the database
migrate-up:
	@migrate -path migrations -database "sqlite3://config/sqlite.db" up
	@echo "Database migrated up."

migrate-down:
	@migrate -path migrations -database "sqlite3://config/sqlite.db" down
	@echo "Database migrated down."

# Target to generate SQLC code
generate-sqlc:
	@sqlc generate
	@echo "SQLC code generated."

.PHONY: run stop log install migrate-up migrate-down generate-sqlc
