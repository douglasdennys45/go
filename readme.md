- cmd/
  - app/
    - main.go
- internal/
  - domain/
    - model/
      - user.go
    - repository/
      - user_repository.go
  - delivery/
    - http/
      - user_handler.go
      - server.go
    - grpc/
      - user_handler.go
      - server.go
- pkg/
  - config/
    - config.go
- go.mod
- go.sum
