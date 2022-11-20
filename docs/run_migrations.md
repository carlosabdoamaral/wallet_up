migrate -path ./migrations -database "postgresql://localhost:5432/wallet_up?sslmode=disable" -verbose up
migrate -path ./migrations -database "postgresql://localhost:5432/wallet_up?sslmode=disable" -verbose down


# CREATE MIGRATION
migrate create -ext sql -dir ./internal/db/migrations -seq init_data