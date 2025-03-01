package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCompetition(t *testing.T) {
	ctx := context.Background()

	newCompetition := CreateCompetitionParams{
		Name:          "Comp 1",
		StartMatchday: 1,
		Status:        NullCompStatus{CompStatus: "open", Valid: true},
	}

	createdCompetition, err := testQueries.CreateCompetition(ctx, newCompetition)

	require.NoError(t, err)
	require.Equal(t, createdCompetition.Name, newCompetition.Name)
	require.Equal(t, createdCompetition.StartMatchday, newCompetition.StartMatchday)
	require.Equal(t, createdCompetition.Status, newCompetition.Status)
}

// TODO: Get competition test
