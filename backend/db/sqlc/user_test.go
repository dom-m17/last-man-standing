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
	ctx := context.Background()

	newUser := CreateUserParams{
		Username:       fmt.Sprintf("user_%s", util.RandomString(5)),
		HashedPassword: fmt.Sprintf("password_%s", util.RandomString(5)),
		FirstName:      util.RandomString(5),
		LastName:       util.RandomString(5),
		Email:          util.RandomEmail(),
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
