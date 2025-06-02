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
			ID:        dbMatch.HomeTeamID,
			LongName:  dbMatch.HomeTeamLongName,
			ShortName: dbMatch.HomeTeamShortName,
			Tla:       dbMatch.HomeTeamTla,
			CrestURL:  dbMatch.HomeTeamCrestUrl.String,
		},
		AwayTeam: models.Team{
			ID:        dbMatch.AwayTeamID,
			LongName:  dbMatch.AwayTeamLongName,
			ShortName: dbMatch.AwayTeamShortName,
			Tla:       dbMatch.AwayTeamTla,
			CrestURL:  dbMatch.AwayTeamCrestUrl.String,
		},
		Matchday:    dbMatch.MatchDate.Day(),
		MatchDate:   dbMatch.MatchDate,
		HomeGoals:   int(dbMatch.HomeGoals.Int32),
		AwayGoals:   int(dbMatch.AwayGoals.Int32),
		HasFinished: dbMatch.HasFinished,
	}
}
