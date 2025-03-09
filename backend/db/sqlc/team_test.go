package db

import (
	"context"
	"database/sql"
	"fmt"
	"lms/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestTeam(t, q)
}

func TestGetTeam(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	newTeam := createTestTeam(t, q)
	foundTeam, err := q.GetTeam(context.Background(), newTeam.ID)

	require.NoError(t, err)
	require.Equal(t, foundTeam.ID, newTeam.ID)
	require.Equal(t, foundTeam.LongName, newTeam.LongName)
	require.Equal(t, foundTeam.ShortName, newTeam.ShortName)
	require.Equal(t, foundTeam.Tla, newTeam.Tla)
	require.Equal(t, foundTeam.CrestUrl, newTeam.CrestUrl)
}

// TODO: List teams is a bit more complicated because it gets all teams from DB so it doesn't work if there are already teams there
// I want the test to work even if there are already existing teams

func createTestTeam(t *testing.T, q *Queries) Team {
	ctx := context.Background()

	teamName := util.RandomString(6)

	newTeam := CreateTeamParams{
		ID:        util.RandomInt(1, 99999),
		LongName:  fmt.Sprintf("%s Town FC", teamName),
		ShortName: teamName,
		Tla:       util.RandomString(3),
		CrestUrl:  sql.NullString{String: fmt.Sprintf("%s.com", teamName), Valid: true},
	}

	createdTeam, err := q.CreateTeam(ctx, newTeam)

	require.NoError(t, err)
	require.Equal(t, createdTeam.ID, newTeam.ID)
	require.Equal(t, createdTeam.LongName, newTeam.LongName)
	require.Equal(t, createdTeam.ShortName, newTeam.ShortName)
	require.Equal(t, createdTeam.Tla, newTeam.Tla)
	require.Equal(t, createdTeam.CrestUrl, newTeam.CrestUrl)

	return createdTeam
}
