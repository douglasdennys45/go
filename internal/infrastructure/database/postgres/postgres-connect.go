package postgres

import (
	"context"
	"database/sql"
	gopostgres "github.com/douglasdennys45/go-postgres"
	"github.com/douglasdennys45/go/internal/infrastructure/controllers"
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

func NewPostgresConnect(driver, dsn string, pool int) error {
	conn := gopostgres.NewPostgresConnection()
	return conn.Connect(driver, dsn, pool)
}

func (m *PostgresConnect) OpenTransaction(ctx context.Context) (*sql.Tx, error) {
	transaction, err := gopostgres.NewPostgresTransaction().OpenTransaction(ctx)
	return transaction, err
}

func (m *PostgresConnect) Commit(tx *sql.Tx) error {
	return gopostgres.NewPostgresTransaction().GetTX().Commit()
}

func (m *PostgresConnect) Rollback(tx *sql.Tx) error {
	return gopostgres.NewPostgresTransaction().GetTX().Rollback()
}

func (m *PostgresConnect) GetTX() *sql.Tx {
	return gopostgres.NewPostgresTransaction().GetTX()
}

func (m *PostgresConnect) GetContext() context.Context {
	return gopostgres.NewPostgresTransaction().GetContext()
}
