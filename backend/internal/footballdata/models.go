package footballdata

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dom-m17/lms/backend/internal/db"
)

type (
	FootballDataMatches struct {
		Matches []Match `json:"matches"`
	}

	Match struct {
		ID        int       `json:"id"`
		HomeTeam  Team      `json:"homeTeam"`
		AwayTeam  Team      `json:"awayTeam"`
		Matchday  int       `json:"matchday"`
		MatchDate time.Time `json:"utcDate"`
	}

	FootballDataTeams struct {
		Teams []Team `json:"teams"`
	}

	Team struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Tla       string `json:"tla"`
		CrestUrl  string `json:"crest"`
	}
)

func makeCreateTeamParams(team Team) db.CreateTeamParams {
	return db.CreateTeamParams{
		ID:        fmt.Sprintf("%d", team.ID),
		LongName:  team.Name,
		ShortName: team.ShortName,
		Tla:       team.Tla,
		CrestUrl:  sql.NullString{String: team.CrestUrl, Valid: team.CrestUrl != ""},
	}
}

func makeCreateMatchParams(match Match) db.CreateMatchParams {
	return db.CreateMatchParams{
		ID:         fmt.Sprintf("%d", match.ID),
		HomeTeamID: fmt.Sprintf("%d", match.HomeTeam.ID),
		AwayTeamID: fmt.Sprintf("%d", match.AwayTeam.ID),
		Matchday:   int32(match.Matchday),
		MatchDate:  match.MatchDate,
	}
}
