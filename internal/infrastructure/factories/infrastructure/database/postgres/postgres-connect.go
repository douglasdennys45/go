package postgres

import (
	"database/sql"

	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	"douglasdenny45.github.com/go/internal/infrastructure/database/postgres"
)

func MountPostgresConnect() *sql.DB {
	return postgres.DB
}

func MountPostgreSQLTransaction() controllers.DBTransaction {
	return postgres.NewPostgreSQLConnect()
}
