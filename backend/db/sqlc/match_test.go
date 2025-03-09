package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateMatch(t *testing.T) {
	q, cleanup := setupTestTransaction(t)
	defer cleanup()

	createTestMatch(t, q)
}

// TODO: Get, List, update match tests

func createTestMatch(t *testing.T, q *Queries) Match {
	ctx := context.Background()

	team1 := createTestTeam(t, q)
	team2 := createTestTeam(t, q)

	newMatch := CreateMatchParams{
		ID:       6,
		HomeTeam: int64(team1.ID),
		// Currently no safeguarding against homeTeam and awayTeam having the same ID
		AwayTeam:  int64(team2.ID),
		Matchday:  1,
		MatchDate: time.Now().UTC(),
	}

	createdMatch, err := q.CreateMatch(ctx, newMatch)

	require.NoError(t, err)
	require.Equal(t, createdMatch.ID, newMatch.ID)
	require.Equal(t, createdMatch.HomeTeam, newMatch.HomeTeam)
	require.Equal(t, createdMatch.AwayTeam, newMatch.AwayTeam)
	require.Equal(t, createdMatch.Matchday, newMatch.Matchday)
	// Create match sends back the time without a timezone so UTC is needed to get the formats to match
	require.Equal(t, createdMatch.MatchDate.UTC(), newMatch.MatchDate)

	return createdMatch
}
