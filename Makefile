run-app:
	go run cmd/main.go
create-migrations:
	migrate create -ext sql -dir database/migrations $(name)
migrations-up:
	migrate -database pgx5://postgres:postgres@localhost:5432/govibes -path database/migrations -verbose up
migrations-down:
	migrate -database pgx5://postgres:postgres@localhost:5432/govibes -path database/migrations -verbose down