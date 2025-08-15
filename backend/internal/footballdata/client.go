package footballdata

import (
	"context"
	"fmt"
	"net/http"
)

const (
	MatchesPath = "/matches"
	TeamsPath   = "/teams"
)

func (s *Service) send(
	ctx context.Context,
	method,
	url string,
) (_ *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating new request with context: %w", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Auth-Token", s.config.FootballDataAPIKey)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending http request: %w", err)
	}

	return res, nil
}

func (s *Service) makeURL(path string) string {
	u := s.config.FootballDataBaseURL.JoinPath(path)

	return u.String()
}
