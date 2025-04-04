// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: match.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMatch = `-- name: CreateMatch :one
INSERT INTO matches (
    id,
  home_team, 
  away_team,
  matchday,
  match_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, home_team, away_team, matchday, match_date, home_goals, away_goals, has_finished
`

type CreateMatchParams struct {
	ID        int64     `json:"id"`
	HomeTeam  int64     `json:"home_team"`
	AwayTeam  int64     `json:"away_team"`
	Matchday  int32     `json:"matchday"`
	MatchDate time.Time `json:"match_date"`
}

func (q *Queries) CreateMatch(ctx context.Context, arg CreateMatchParams) (Match, error) {
	row := q.db.QueryRowContext(ctx, createMatch,
		arg.ID,
		arg.HomeTeam,
		arg.AwayTeam,
		arg.Matchday,
		arg.MatchDate,
	)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.HomeTeam,
		&i.AwayTeam,
		&i.Matchday,
		&i.MatchDate,
		&i.HomeGoals,
		&i.AwayGoals,
		&i.HasFinished,
	)
	return i, err
}

const getMatch = `-- name: GetMatch :one
SELECT id, home_team, away_team, matchday, match_date, home_goals, away_goals, has_finished FROM matches
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMatch(ctx context.Context, id int64) (Match, error) {
	row := q.db.QueryRowContext(ctx, getMatch, id)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.HomeTeam,
		&i.AwayTeam,
		&i.Matchday,
		&i.MatchDate,
		&i.HomeGoals,
		&i.AwayGoals,
		&i.HasFinished,
	)
	return i, err
}

const listMatches = `-- name: ListMatches :many
SELECT id, home_team, away_team, matchday, match_date, home_goals, away_goals, has_finished FROM matches
WHERE matchday = $1
ORDER BY id
`

func (q *Queries) ListMatches(ctx context.Context, matchday int32) ([]Match, error) {
	rows, err := q.db.QueryContext(ctx, listMatches, matchday)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Match{}
	for rows.Next() {
		var i Match
		if err := rows.Scan(
			&i.ID,
			&i.HomeTeam,
			&i.AwayTeam,
			&i.Matchday,
			&i.MatchDate,
			&i.HomeGoals,
			&i.AwayGoals,
			&i.HasFinished,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMatch = `-- name: UpdateMatch :one
UPDATE matches
SET home_goals = $1,
 away_goals = $2
WHERE id = $3
RETURNING id, home_team, away_team, matchday, match_date, home_goals, away_goals, has_finished
`

type UpdateMatchParams struct {
	HomeGoals sql.NullInt32 `json:"home_goals"`
	AwayGoals sql.NullInt32 `json:"away_goals"`
	ID        int64         `json:"id"`
}

func (q *Queries) UpdateMatch(ctx context.Context, arg UpdateMatchParams) (Match, error) {
	row := q.db.QueryRowContext(ctx, updateMatch, arg.HomeGoals, arg.AwayGoals, arg.ID)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.HomeTeam,
		&i.AwayTeam,
		&i.Matchday,
		&i.MatchDate,
		&i.HomeGoals,
		&i.AwayGoals,
		&i.HasFinished,
	)
	return i, err
}
