package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func NewMysqlConnect(driver, dsn string, pool int) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(1000)
	db.SetMaxOpenConns(pool)
	db.SetMaxIdleConns(pool)
	DB = db
	return db, nil
}

func Begin() (*sql.Tx, error) {
	ctx, err := DB.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}
