package user

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	token "github.com/dom-m17/lms/backend/internal/refresh-token"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

type Service struct {
	Querier      db.Querier
	RefreshToken *token.Service
}

type ServiceInterface interface {
	GetUser(ctx context.Context, input string) (*models.User, error)
	ListUsers(ctx context.Context) (models.Users, error)
	CreateUser(ctx context.Context, input graphmodels.CreateUserInput) (*models.User, error)
	DeleteUser(ctx context.Context, userID string) (*models.User, error)
	UpdateUser(ctx context.Context, input graphmodels.UpdateUserInput) (*models.User, error)
	Login(ctx context.Context, input graphmodels.LoginInput) (*models.User, string, string, error)
}

func NewService(querier db.Querier, token *token.Service) *Service {
	return &Service{
		Querier:      querier,
		RefreshToken: token,
	}
}
