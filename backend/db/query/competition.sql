-- name: CreateCompetition :one
INSERT INTO competition (
  name, 
  start_matchday, 
  status,
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCompetition :one
SELECT * FROM competition
WHERE id = $1 LIMIT 1;
