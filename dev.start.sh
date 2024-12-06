#!/bin/zsh


echo "Start database"
alacritty -e bash -c "docker compose -f docker-compose.dev.yml up" &


echo "Start backend"
alacritty -e bash -c "cd backend && go run cmd/main.go" &


echo "Start frontend"
bash -c "cd frontend && npm run dev"
