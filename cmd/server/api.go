package main

import (
	"github.com/douglasdennys45/go/internal/infrastructure/config"
	"github.com/douglasdennys45/go/internal/infrastructure/database/postgres"
	"github.com/douglasdennys45/go/pkg/env"
	"github.com/douglasdennys45/go/pkg/logger"
)

func init() {
	env, err := env.NewEnv(".")
	if err != nil {
		panic("env is nil")
	}
	err = postgres.NewPostgresConnect(env.DBDriver, env.DBDSN, env.DBPool)
	if err != nil {
		panic(err)
	}
}

func main() {
	log := logger.NewLogger()
	app := config.NewServer()
	log.Info("Server running on port 3000")
	app.Listen(":3000")
}
