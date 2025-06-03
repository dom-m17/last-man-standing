package competition

import (
	"context"
	"fmt"

	"github.com/dom-m17/lms/backend/internal/models"
)

func (s *Service) GetCompetition(ctx context.Context, competitionID string) (models.Competition, error) {
	competition, err := s.Querier.GetCompetition(ctx, competitionID)
	if err != nil {
		return models.Competition{}, fmt.Errorf("getting competition: %w", err)
	}

	return convertDBCompetitionToModelsCompetition(competition), nil
}
