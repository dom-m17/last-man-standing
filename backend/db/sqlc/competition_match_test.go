package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCompetitionMatch(t *testing.T) {
	ctx := context.Background()

	newCompetitionMatch := CreateCompetitionMatchParams{
		CompetitionID: 1,
		MatchID:       1,
	}

	createdCompetitionMatch, err := testQueries.CreateCompetitionMatch(ctx, newCompetitionMatch)

	require.NoError(t, err)
	require.Equal(t, createdCompetitionMatch.CompetitionID, newCompetitionMatch.CompetitionID)
	require.Equal(t, createdCompetitionMatch.MatchID, newCompetitionMatch.MatchID)
}

// TODO: List competition matches tests
