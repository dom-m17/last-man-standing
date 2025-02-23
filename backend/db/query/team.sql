-- name: CreateTeam :one
INSERT INTO teams (
  id,
  long_name, 
  short_name, 
  tla, 
  crest_url
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTeam :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM teams
LIMIT 20;