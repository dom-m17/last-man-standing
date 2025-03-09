package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/lms?sslmode=disable"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	os.Exit(m.Run())
}

func setupTestTransaction(t *testing.T) (*Queries, func()) {
	tx, err := testDB.Begin()
	if err != nil {
		t.Fatalf("failed to start transaction: %v", err)
	}

	q := New(tx)

	cleanup := func() {
		if err := tx.Rollback(); err != nil {
			t.Fatalf("failed to rollback transaction: %v", err)
		}
	}

	return q, cleanup
}
