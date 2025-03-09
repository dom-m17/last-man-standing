package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCompetition(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestCompetition(t, q)
}

// TODO: Get competition test

func createTestCompetition(t *testing.T, q *Queries) Competition {
	ctx := context.Background()

	newCompetition := CreateCompetitionParams{
		Name:          "Comp 1",
		StartMatchday: 1,
		Status:        NullCompStatus{CompStatus: "open", Valid: true},
	}

	createdCompetition, err := q.CreateCompetition(ctx, newCompetition)

	require.NoError(t, err)
	require.Equal(t, createdCompetition.Name, newCompetition.Name)
	require.Equal(t, createdCompetition.StartMatchday, newCompetition.StartMatchday)
	require.Equal(t, createdCompetition.Status, newCompetition.Status)

	return createdCompetition
}
