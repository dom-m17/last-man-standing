package competition

import (
	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

func convertDBCompetitionToModelsCompetition(competition db.Competition) models.Competition {
	return models.Competition{
		ID:            competition.ID,
		Name:          competition.Name,
		StartMatchday: int(competition.StartMatchday),
		Status:        competition.Status,
	}
}
