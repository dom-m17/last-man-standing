package competition

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

type Service struct {
	Querier db.Querier
}

type ServiceInterface interface {
	GetCompetition(ctx context.Context, competitionID string) (models.Competition, error)
}

func NewService(querier db.Querier) *Service {
	return &Service{
		Querier: querier,
	}
}
