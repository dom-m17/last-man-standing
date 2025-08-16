package footballdata

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
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
		Score     Score     `json:"score"`
		Status    string    `json:"status"`
	}

	Score struct {
		FullTime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"fullTime"`
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

func makeCreateMatchParams(match Match) db.CreateUpdateMatchParams {
	return db.CreateUpdateMatchParams{
		ID:         fmt.Sprintf("%d", match.ID),
		HomeTeamID: fmt.Sprintf("%d", match.HomeTeam.ID),
		AwayTeamID: fmt.Sprintf("%d", match.AwayTeam.ID),
		Matchday:   int32(match.Matchday),
		MatchDate:  match.MatchDate,
		HomeGoals:  convertAPIGoalsToDBGoals(match.Status, match.Score.FullTime.Home),
		AwayGoals:  convertAPIGoalsToDBGoals(match.Status, match.Score.FullTime.Away),
		// The statuses used by the API are the same as the ones used in the DB
		Status: models.MatchStatus(match.Status),
	}
}

func convertAPIGoalsToDBGoals(status string, goals int) sql.NullInt32 {
	if status == string(models.MatchStatusInProgress) || status == string(models.MatchStatusFinished) {
		return sql.NullInt32{Int32: int32(goals), Valid: true}
	}

	return sql.NullInt32{Int32: 0, Valid: false}
}
