package selection

import (
	"github.com/dom-m17/lms/backend/internal/db"
)

type Service struct {
	Querier db.Querier
}

type ServiceInterface interface {
}

func NewService(querier db.Querier) *Service {
	return &Service{
		Querier: querier,
	}
}
