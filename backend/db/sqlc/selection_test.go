package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSelection(t *testing.T) {
	ctx := context.Background()

	newSelection := CreateSelectionParams{
		EntryID: 1,
		MatchID: 1,
		TeamID:  1,
	}

	createdMatch, err := testQueries.CreateSelection(ctx, newSelection)

	require.NoError(t, err)
	require.Equal(t, createdMatch.EntryID, newSelection.EntryID)
	require.Equal(t, createdMatch.MatchID, newSelection.MatchID)
	require.Equal(t, createdMatch.TeamID, newSelection.TeamID)
}

// TODO: Get, List selection tests
