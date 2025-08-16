package match

import (
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
		ID: dbMatch.MatchID,
		HomeTeam: models.Team{
			ID: dbMatch.HomeTeamID,
		},
		AwayTeam: models.Team{
			ID: dbMatch.AwayTeamID,
		},
		Matchday:  int(dbMatch.Matchday),
		MatchDate: dbMatch.MatchDate,
		HomeGoals: int(dbMatch.HomeGoals.Int32),
		AwayGoals: int(dbMatch.AwayGoals.Int32),
		Status:    models.MatchStatus(dbMatch.Status),
	}
}
