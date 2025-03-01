package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	newUser := CreateUserParams{
		Username:       "dom_m17",
		HashedPassword: "password",
		FirstName:      "Dominic",
		LastName:       "Maynard",
		Email:          "example@example.com",
		PhoneNumber:    sql.NullString{String: "07123456789", Valid: true},
		FavouriteTeam:  sql.NullInt64{Int64: 1, Valid: true},
	}

	user, err := testQueries.CreateUser(ctx, newUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, newUser.Username)
	require.Equal(t, user.HashedPassword, newUser.HashedPassword)
	require.Equal(t, user.FirstName, newUser.FirstName)
	require.Equal(t, user.LastName, newUser.LastName)
	require.Equal(t, user.Email, newUser.Email)
	require.Equal(t, user.PhoneNumber, newUser.PhoneNumber)
	require.Equal(t, user.FavouriteTeam, newUser.FavouriteTeam)
}

// TODO: Get User test
