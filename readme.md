mockgen -source=internal/domain/repositories/user-repo.go -destination=internal/domain/repositories/mocks/mock-user-repo.go -package=mocks

go test ./...


go test -cover ./... && go test -cover -coverprofile=coverage.out ./... && go tool cover -html=c.out -o coverage.html