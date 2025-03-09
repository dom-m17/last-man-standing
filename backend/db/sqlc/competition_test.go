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

func TestGetCompetition(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	newCompetition := createTestCompetition(t, q)
	foundCompetition, err := q.GetCompetition(context.Background(), newCompetition.ID)

	require.NoError(t, err)
	require.Equal(t, foundCompetition.Name, newCompetition.Name)
	require.Equal(t, foundCompetition.StartMatchday, newCompetition.StartMatchday)
	require.Equal(t, foundCompetition.Status, newCompetition.Status)
}

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
