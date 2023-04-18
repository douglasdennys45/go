package postgres_test

import (
	"github.com/douglasdennys45/go/internal/infrastructure/database/postgres"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupDB() {
	database := embeddedpostgres.NewDatabase()
	_ = database.Start()
	defer database.Stop()
}

func TestNewPostgresConnect(t *testing.T) {
	setupDB()
	err := postgres.NewPostgresConnect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", 10)
	assert.Nil(t, err)
}

/*func TestNewPostgreSQLConnectOpenTransaction(t *testing.T) {
	setupDB()
	_ = postgres.NewPostgresConnect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", 10)
	db := postgres.NewPostgreSQLConnect()
	tx, err := db.OpenTransaction(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, tx)
}

func TestNewPostgreSQLConnectCommit(t *testing.T) {
	setupDB()
	_ = postgres.NewPostgresConnect("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", 10)
	db := postgres.NewPostgreSQLConnect()
	_, _ = db.OpenTransaction(context.Background())
	err := db.Commit(db.GetTX())
	assert.Nil(t, err)
}*/
