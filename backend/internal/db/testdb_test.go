package db

import (
	"testing"

	"github.com/peterldowns/testy/check"
)

func TestAQuery(t *testing.T) {
	t.Parallel()
	db := NewDB(t)
	defer db.Close()

	var result string
	err := db.QueryRow("SELECT 'hello world'").Scan(&result)
	check.Nil(t, err)
	check.Equal(t, "hello world", result)
}
