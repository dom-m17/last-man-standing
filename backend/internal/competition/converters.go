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
		Status:        convertDBCompStatusToCompStatus(competition.Status),
	}
}

func convertDBCompStatusToCompStatus(status db.CompStatus) models.CompStatus {
	switch status {
	case db.CompStatusOpen:
		return models.CompStatusOpen
	case db.CompStatusInProgress:
		return models.CompStatusInProgress
	case db.CompStatusComplete:
		return models.CompStatusComplete
	}

	// TODO: Handle this better. I am not sure what the idiomatic way is
	return models.CompStatus("")
}
