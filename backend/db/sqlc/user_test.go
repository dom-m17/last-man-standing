package db

import (
	"context"
	"database/sql"
	"fmt"
	"lms/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestUser(t, q)
}

// TODO: Get User test

func createTestUser(t *testing.T, q *Queries) User {
	ctx := context.Background()

	team := createTestTeam(t, q)

	newUser := CreateUserParams{
		Username:       fmt.Sprintf("user_%s", util.RandomString(5)),
		HashedPassword: fmt.Sprintf("password_%s", util.RandomString(5)),
		FirstName:      util.RandomString(5),
		LastName:       util.RandomString(5),
		Email:          util.RandomEmail(),
		PhoneNumber:    sql.NullString{String: "07123456789", Valid: true},
		FavouriteTeam:  sql.NullInt64{Int64: int64(team.ID), Valid: true},
	}

	createdUser, err := q.CreateUser(ctx, newUser)

	require.NoError(t, err)
	require.Equal(t, createdUser.Username, newUser.Username)
	require.Equal(t, createdUser.HashedPassword, newUser.HashedPassword)
	require.Equal(t, createdUser.FirstName, newUser.FirstName)
	require.Equal(t, createdUser.LastName, newUser.LastName)
	require.Equal(t, createdUser.Email, newUser.Email)
	require.Equal(t, createdUser.PhoneNumber, newUser.PhoneNumber)
	require.Equal(t, createdUser.FavouriteTeam, newUser.FavouriteTeam)

	return createdUser
}
