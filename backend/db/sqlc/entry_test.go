package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	ctx := context.Background()

	newEntry := CreateEntryParams{
		UserID:        3,
		CompetitionID: 1,
	}

	createdEntry, err := testQueries.CreateEntry(ctx, newEntry)

	require.NoError(t, err)
	require.Equal(t, createdEntry.UserID, newEntry.UserID)
	// No competition yet so won't work
	require.Equal(t, createdEntry.CompetitionID, newEntry.CompetitionID)
}

// TODO: Get and List Entry tests
