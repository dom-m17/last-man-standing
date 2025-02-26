-- name: CreateCompetitionMatch :one
INSERT INTO competition_matches (
  competition_id,
  match_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListCompetitionMatches :many
SELECT * FROM competition_matches
WHERE competition_id = $1
ORDER BY competition_id;