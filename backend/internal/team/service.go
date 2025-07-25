package team

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

type Service struct {
	Querier db.Querier
}

type ServiceInterface interface {
	GetTeam(ctx context.Context, teamID string) (models.Team, error)
	ListTeams(ctx context.Context) (models.Teams, error)
}

func NewService(querier db.Querier) *Service {
	return &Service{
		Querier: querier,
	}
}
