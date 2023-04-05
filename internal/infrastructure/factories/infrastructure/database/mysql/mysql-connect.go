package mysql

import (
	"database/sql"

	"douglasdenny45.github.com/go/internal/infrastructure/database/mysql"
)

func MountMySQLConnect() *sql.DB {
	return mysql.DB
}
