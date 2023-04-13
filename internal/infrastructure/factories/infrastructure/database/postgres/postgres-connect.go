package postgres

import (
	"database/sql"

	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
	"github.com/douglasdennys45/go/internal/infrastructure/database/postgres"
)

func MountPostgresConnect() *sql.DB {
	return postgres.DB
}

func MountPostgreSQLTransaction() controllers.DBTransaction {
	return postgres.NewPostgreSQLConnect()
}
