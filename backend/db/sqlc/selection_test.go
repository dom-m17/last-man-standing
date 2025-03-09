package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSelection(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestSelection(t, q)
}

// TODO: Get, List selection tests

func createTestSelection(t *testing.T, q *Queries) {
	ctx := context.Background()

	match := createTestMatch(t, q)
	entry := createTestEntry(t, q)

	newSelection := CreateSelectionParams{
		EntryID: entry.ID,
		MatchID: match.ID,
		TeamID:  match.HomeTeam,
	}

	createdSelection, err := q.CreateSelection(ctx, newSelection)

	require.NoError(t, err)
	require.Equal(t, createdSelection.EntryID, newSelection.EntryID)
	require.Equal(t, createdSelection.MatchID, newSelection.MatchID)
	require.Equal(t, createdSelection.TeamID, newSelection.TeamID)
}
