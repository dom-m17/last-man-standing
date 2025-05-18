package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/peterldowns/testy/check"
)

// This is a helper function because any other test will require creating a user first
func createTestUser(t *testing.T, ctx context.Context, q Querier) User {

	//TODO: Use random parameters so this can be used to create multiple users
	createdUser, err := q.CreateUser(ctx, CreateUserParams{
		Username:       "dom_m17",
		HashedPassword: "password",
		FirstName:      "Dominic",
		LastName:       "Maynard",
		Email:          "dom@email.com",
		PhoneNumber:    sql.NullString{String: "0123", Valid: true},
	})
	check.Nil(t, err)
	check.Equal(t, createdUser.Username, "dom_m17")
	check.Equal(t, createdUser.HashedPassword, "password")
	check.Equal(t, createdUser.FirstName, "Dominic")
	check.Equal(t, createdUser.LastName, "Maynard")
	check.Equal(t, createdUser.Email, "dom@email.com")
	check.Equal(t, createdUser.PhoneNumber.String, "0123")

	return createdUser
}

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	q, close := NewTestQuerier(t)
	defer close()
	createTestUser(t, ctx, q)
}

func Test_GetUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	q, close := NewTestQuerier(t)
	defer close()

	createdUser := createTestUser(t, ctx, q)
	gotUser, err := q.GetUser(ctx, createdUser.ID)
	check.Nil(t, err)
	check.Equal(t, createdUser, gotUser)
}
