package match

import (
	"fmt"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

func convertDBMatchesToModelsMatches(dbMatches []db.GetMatchesByMatchdayRow) []*models.Match {
	modelMatches := make([]*models.Match, len(dbMatches))
	for i := range dbMatches {
		modelMatches[i] = convertDBMatchToModelsMatch(dbMatches[i])
	}

	return modelMatches
}

func convertDBMatchToModelsMatch(dbMatch db.GetMatchesByMatchdayRow) *models.Match {
	return &models.Match{
		ID:           fmt.Sprintf("%d", dbMatch.MatchID),
		HomeTeamID:   int(dbMatch.HomeTeam),
		HomeTeamName: dbMatch.HomeTeamName,
		AwayTeamID:   int(dbMatch.AwayTeam),
		AwayTeamName: dbMatch.AwayTeamName,
		Matchday:     dbMatch.MatchDate.Day(),
		MatchDate:    dbMatch.MatchDate,
		HomeGoals:    int(dbMatch.HomeGoals.Int32),
		AwayGoals:    int(dbMatch.AwayGoals.Int32),
		HasFinished:  dbMatch.HasFinished,
	}
}
