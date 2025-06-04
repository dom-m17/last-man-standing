package competition

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

type Service struct {
	Querier db.Querier
}

type ServiceInterface interface {
	GetCompetition(ctx context.Context, competitionID string) (models.Competition, error)
	CreateCompetition(ctx context.Context, input graphmodels.CreateCompetitionInput) (models.Competition, error)
}

func NewService(querier db.Querier) *Service {
	return &Service{
		Querier: querier,
	}
}
