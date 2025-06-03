package team

import (
	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

func convertDBTeamsToModelsTeams(teams []db.Team) models.Teams {
	mTeams := make(models.Teams, len(teams))
	for i := range teams {
		mTeam := convertDBTeamToModelsTeam(teams[i])
		mTeams[i] = &mTeam
	}

	return mTeams
}

func convertDBTeamToModelsTeam(team db.Team) models.Team {
	return models.Team{
		ID:        team.ID,
		LongName:  team.LongName,
		ShortName: team.ShortName,
		Tla:       team.Tla,
		CrestURL:  team.CrestUrl.String,
	}
}
