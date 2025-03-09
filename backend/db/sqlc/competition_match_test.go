package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCompetitionMatch(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestCompetitionMatch(t, q)
}

// TODO: List competition matches tests

func createTestCompetitionMatch(t *testing.T, q *Queries) {
	ctx := context.Background()

	comp := createTestCompetition(t, q)
	match := createTestMatch(t, q)

	newCompetitionMatch := CreateCompetitionMatchParams{
		CompetitionID: comp.ID,
		MatchID:       match.ID,
	}

	createdCompetitionMatch, err := q.CreateCompetitionMatch(ctx, newCompetitionMatch)

	require.NoError(t, err)
	require.Equal(t, createdCompetitionMatch.CompetitionID, newCompetitionMatch.CompetitionID)
	require.Equal(t, createdCompetitionMatch.MatchID, newCompetitionMatch.MatchID)
}
