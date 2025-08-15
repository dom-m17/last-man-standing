package footballdata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *Service) PopulateTeams(ctx context.Context) error {
	url := s.makeURL(TeamsPath)

	res, err := s.send(ctx, http.MethodGet, url)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %d", "unexpected status code", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading body: %w", err)
	}

	var data FootballDataTeams
	err = json.Unmarshal(body, &data)
	if err != nil {
		return fmt.Errorf("unmarshaling: %w", err)
	}

	// TODO: find out if there is a way to insert many at one time with sqlc
	for _, team := range data.Teams {
		s.Querier.CreateTeam(ctx, makeCreateTeamParams(team))
	}

	return nil
}
