package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateMatch(t *testing.T) {
	ctx := context.Background()

	newMatch := CreateMatchParams{
		ID:        6,
		HomeTeam:  1,
		AwayTeam:  1,
		Matchday:  1,
		MatchDate: time.Now().UTC(),
	}

	createdMatch, err := testQueries.CreateMatch(ctx, newMatch)

	require.NoError(t, err)
	require.Equal(t, createdMatch.ID, newMatch.ID)
	require.Equal(t, createdMatch.HomeTeam, newMatch.HomeTeam)
	require.Equal(t, createdMatch.AwayTeam, newMatch.AwayTeam)
	require.Equal(t, createdMatch.Matchday, newMatch.Matchday)
	// Create match sends back the time without a timezone so UTC is needed to get the formats to match
	require.Equal(t, createdMatch.MatchDate.UTC(), newMatch.MatchDate)
}

// TODO: Get, List, update match tests
