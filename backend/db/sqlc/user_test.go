package db

import (
	"database/sql"
	"testing"

	"github.com/peterldowns/testy/check"
)

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	db := NewDB(t)
	defer db.Close()
	q := New(db)

	createdUser, err := q.CreateUser(ctx, CreateUserParams{
		Username:       "dom_m17",
		HashedPassword: "password",
		FirstName:      "Dominic",
		LastName:       "Maynard",
		Email:          "dom@email.com",
		PhoneNumber:    sql.NullString{String: "0123", Valid: true},
	})
	check.Nil(t, err)
	gotUser, err := q.GetUser(ctx, createdUser.ID)
	check.Nil(t, err)
	check.Equal(t, createdUser, gotUser)
}
