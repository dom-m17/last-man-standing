package competition

import (
	"context"
	"fmt"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

func (s *Service) GetCompetition(ctx context.Context, competitionID string) (models.Competition, error) {
	competition, err := s.Querier.GetCompetition(ctx, competitionID)
	if err != nil {
		return models.Competition{}, fmt.Errorf("getting competition: %w", err)
	}

	return convertDBCompetitionToModelsCompetition(competition), nil
}

func (s *Service) ListCompetitions(ctx context.Context) (models.Competitions, error) {
	competitions, err := s.Querier.ListCompetitions(ctx)
	if err != nil {
		return models.Competitions{}, fmt.Errorf("listing competitions: %w", err)
	}

	return convertDBCompetitionsToModelsCompetitions(competitions), nil
}

func (s *Service) CreateCompetition(ctx context.Context, input graphmodels.CreateCompetitionInput) (models.Competition, error) {
	competition, err := s.Querier.CreateCompetition(ctx, db.CreateCompetitionParams{
		Name:          input.Name,
		StartMatchday: input.StartMatchday,
	})
	if err != nil {
		return models.Competition{}, fmt.Errorf("creating competition: %w", err)
	}

	return convertDBCompetitionToModelsCompetition(competition), nil
}
