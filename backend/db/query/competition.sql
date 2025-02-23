-- name: CreateCompetition :one
INSERT INTO competitions (
  name, 
  start_matchday, 
  status
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCompetition :one
SELECT * FROM competitions
WHERE id = $1 LIMIT 1;
