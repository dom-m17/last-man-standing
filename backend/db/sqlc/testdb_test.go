package db

import (
	"testing"

	"github.com/peterldowns/pgtestdb"
	"github.com/peterldowns/testy/check"
	"github.com/stretchr/testify/assert"
)

func TestAQuery(t *testing.T) {
	t.Parallel()
	db := NewDB(t) // this is the helper defined above

	var result string
	err := db.QueryRow("SELECT 'hello world'").Scan(&result)
	check.Nil(t, err)
	check.Equal(t, "hello world", result)
}

func TestWithPgxStdlibDriver(t *testing.T) {
	t.Parallel()
	pgxConf := pgtestdb.Config{
		DriverName: "pgx", // uses the pgx/stdlib driver
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		Port:       "5433",
		Options:    "sslmode=disable",
	}
	migrator := pgtestdb.NoopMigrator{}
	db := pgtestdb.New(t, pgxConf, migrator)

	var message string
	err := db.QueryRow("select 'hello world'").Scan(&message)
	assert.Nil(t, err)
	assert.Equal(t, "hello world", message)
}
