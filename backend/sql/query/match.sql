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
    "status" AS "status"
FROM matches
WHERE matchday = $1;

-- name: UpsertMatch :one
INSERT INTO matches (
    id, home_team_id, away_team_id, matchday, match_date, home_goals, away_goals, "status"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) 
ON CONFLICT (id) DO UPDATE
SET home_team_id = EXCLUDED.home_team_id,
    away_team_id = EXCLUDED.away_team_id,
    matchday = EXCLUDED.matchday,
    match_date = EXCLUDED.match_date,
    home_goals = EXCLUDED.home_goals,
    away_goals = EXCLUDED.away_goals,
    "status" = EXCLUDED."status"
RETURNING *;
