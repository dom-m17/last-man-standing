-- name: GetMatch :one
SELECT * FROM matches
WHERE id = $1;

-- name: GetMatchesByMatchday :many
SELECT * FROM matches
WHERE matchday = $1;

-- name: CreateMatch :one
INSERT INTO matches (
    home_team, away_team, matchday, match_date
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