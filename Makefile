.PHONY: all db backend frontend

all: db backend frontend

db:
	@echo "Start database"
	bash -c "docker compose -f docker-compose.dev.yml up" &

backend:
	@echo "Start backend"
	bash -c "cd backend && go run cmd/main.go" &

frontend:
	@echo "Start frontend"
	 bash -c "cd frontend && npm run dev"
