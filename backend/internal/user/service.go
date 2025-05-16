package user

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	"github.com/dom-m17/lms/backend/internal/subgraph/model"
)

type Service struct {
	Querier db.Querier
}

type ServiceInterface interface {
	GetUser(ctx context.Context, input string) (*models.User, error)
	CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error)
}

func NewService(querier db.Querier) *Service {
	return &Service{
		Querier: querier,
	}
}
