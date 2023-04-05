FROM go:1.20-alpine3.13 AS builder
LABEL maintainer="Douglas Dennys <douglasdennys45@gmail.com>"
WORKDIR /app
COPY . .
RUN go mod download && GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/server/api.go


FROM scratch
COPY --from=builder /app/main .
CMD ["./main"]