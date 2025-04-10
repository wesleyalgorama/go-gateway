build:
	go build cmd/app/main.go

start:
	./main

dev:
	go run cmd/app/main.go

migrate: 
	migrate -database "postgres://user:pass@localhost:5432/db?sslmode=disable" -path migrations up