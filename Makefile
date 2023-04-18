test:
	go test ./... `go list ./... | grep -v ./internal/infrastructure/middlewares | grep -v ./internal/infrastructure/factories`

run:
	go run cmd/server/api.go

generate:
	go test `go list ./... | grep -v ./internal/infrastructure/middlewares | grep -v ./internal/infrastructure/factories` -cover -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

coverage:
	open coverage.html

test-bench:
	autocannon -m POST -b '{"name": "John Doe","email": "password","password": "password"}' -H "Content-Type: application/json" -a 500 -c 10 http://localhost:3000/users