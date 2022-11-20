api:
	go run cmd/api/main.go

rabbit:
	go run cmd/rabbit/main.go

grpc:
	go run cmd/grpc/main.go

docker:
	cd docs/ && docker-compose up -d

migrateup:
	migrate -path ./internal/db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/db/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down