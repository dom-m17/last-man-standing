-- name: CreateMatch :one
INSERT INTO matches (
    id,
  home_team, 
  away_team,
  matchday,
  match_date
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetMatch :one
SELECT * FROM matches
WHERE id = $1 LIMIT 1;

-- name: ListMatches :many
SELECT * FROM matches
WHERE matchday = $1
ORDER BY id;

-- name: UpdateMatch :one
UPDATE matches
SET home_goals = $1,
 away_goals = $2
WHERE id = $3
RETURNING *;