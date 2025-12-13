package refreshtoken

import (
	"context"

	"github.com/dom-m17/lms/backend/internal/models"
)

func (s *Service) GetRefreshTokenByTokenHash(ctx context.Context, refreshToken string) (*models.RefreshToken, error) {
	return &models.RefreshToken{}, nil
}

// LMS-45
func (s *Service) CreateRefreshToken() string {
	return ""
}
