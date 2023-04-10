package postgres

import (
	"context"
	"database/sql"

	"douglasdenny45.github.com/go/internal/infrastructure/controllers"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

var (
	TX  *sql.Tx
	CTX context.Context
)

type PostgresConnect struct {
}

func NewPostgreSQLConnect() controllers.DBTransaction {
	return &PostgresConnect{}
}

func NewPostgresConnect(driver, dsn string, pool int) (*sql.DB, error) {
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

func (m *PostgresConnect) OpenTransaction(ctx context.Context) (*sql.Tx, error) {
	tx, err := DB.BeginTx(ctx, nil)
	TX = tx
	CTX = ctx
	return TX, err
}

func (m *PostgresConnect) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (m *PostgresConnect) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func (m *PostgresConnect) GetTX() *sql.Tx {
	return TX
}

func (m *PostgresConnect) GetContext() context.Context {
	return CTX
}
