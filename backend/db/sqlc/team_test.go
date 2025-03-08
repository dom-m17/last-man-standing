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
	defer cleanup() // Ensure transaction is rolled back after the test

	createTestTeam(t, q)
}

func createTestTeam(t *testing.T, q *Queries) {
	ctx := context.Background()

	teamName := util.RandomString(6)

	newTeam := CreateTeamParams{
		ID:        util.RandomInt(100, 999),
		LongName:  fmt.Sprintf("%s Town FC", teamName),
		ShortName: teamName,
		Tla:       util.RandomString(3),
		CrestUrl:  sql.NullString{String: fmt.Sprintf("%s.com", teamName), Valid: true},
	}

	// Use the transaction-bound queries
	createdTeam, err := q.CreateTeam(ctx, newTeam)

	require.NoError(t, err)
	require.Equal(t, newTeam.ID, createdTeam.ID)
	require.Equal(t, newTeam.LongName, createdTeam.LongName)
	require.Equal(t, newTeam.ShortName, createdTeam.ShortName)
	require.Equal(t, newTeam.Tla, createdTeam.Tla)
	require.Equal(t, newTeam.CrestUrl, createdTeam.CrestUrl)
}

// TODO: Get team and List team tests
