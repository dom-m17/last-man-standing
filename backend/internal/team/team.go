package team

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/models"
)

func (s *Service) GetTeam(ctx context.Context, teamID string) (models.Team, error) {
	team, err := s.Querier.GetTeam(ctx, teamID)
	if err != nil {
		return models.Team{}, nil
	}

	return convertDBTeamToModelsTeam(team), nil
}

func (s *Service) ListTeams(ctx context.Context) (models.Teams, error) {
	teams, err := s.Querier.ListTeams(ctx)
	if err != nil {
		return models.Teams{}, nil
	}

	return convertDBTeamsToModelsTeams(teams), nil
}
