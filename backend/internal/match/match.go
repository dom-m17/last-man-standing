package match

import (
	"context"
	"fmt"

	"github.com/dom-m17/lms/backend/internal/models"
)

func (s *Service) GetMatch(ctx context.Context, id string) (*models.Match, error) {
	panic(fmt.Errorf("not implemented: GetMatch - getMatch"))
}

func (s *Service) GetMatchesByMatchday(ctx context.Context, matchday int32) ([]*models.Match, error) {
	matches, err := s.Querier.GetMatchesByMatchday(ctx, matchday)
	if err != nil {
		return nil, fmt.Errorf("getting matches by matchday: %w", err)
	}

	return convertDBMatchesToModelsMatches(matches), nil
}
