package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestEntry(t, q)
}

// TODO: Get and List Entry tests

func createTestEntry(t *testing.T, q *Queries) Entry {
	ctx := context.Background()

	user := createTestUser(t, q)
	competition := createTestCompetition(t, q)

	newEntry := CreateEntryParams{
		UserID:        user.ID,
		CompetitionID: competition.ID,
	}

	createdEntry, err := q.CreateEntry(ctx, newEntry)

	require.NoError(t, err)
	require.Equal(t, createdEntry.UserID, newEntry.UserID)
	// No competition yet so won't work
	require.Equal(t, createdEntry.CompetitionID, newEntry.CompetitionID)

	return createdEntry
}
