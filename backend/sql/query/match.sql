-- name: GetMatch :one
SELECT * FROM matches
WHERE id = $1;

-- name: GetMatchesByMatchday :many
SELECT 
    id AS match_id,
    home_team_id AS home_team_id,
    away_team_id AS away_team_id,
    matchday AS matchday,
    match_date AS match_date,
    home_goals AS home_goals,
    away_goals AS away_goals,
    has_finished AS has_finished
FROM matches
WHERE matchday = $1;

-- name: CreateMatch :one
INSERT INTO matches (
    home_team_id, away_team_id, matchday, match_date
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateMatch :one
UPDATE matches 
SET
    matchday = $2,
    match_date = $3,
    home_goals = $4,
    away_goals = $5,
    has_finished = "TRUE"
WHERE id = $1
RETURNING *;