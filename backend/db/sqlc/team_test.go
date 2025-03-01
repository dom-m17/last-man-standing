package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {
	ctx := context.Background()

	newTeam := CreateTeamParams{
		ID:        1,
		LongName:  "Arsenal FC",
		ShortName: "Arsenal",
		Tla:       "ARS",
		CrestUrl:  sql.NullString{String: "arsenalcrest.com", Valid: true},
	}

	createdTeam, err := testQueries.CreateTeam(ctx, newTeam)

	require.NoError(t, err)

	require.Equal(t, createdTeam.ID, newTeam.ID)
	require.Equal(t, createdTeam.LongName, newTeam.LongName)
	require.Equal(t, createdTeam.ShortName, newTeam.ShortName)
	require.Equal(t, createdTeam.Tla, newTeam.Tla)
	require.Equal(t, createdTeam.CrestUrl, newTeam.CrestUrl)
}

// TODO: Get team and List team tests
