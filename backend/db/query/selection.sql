-- name: CreateSelection :one
INSERT INTO selections (
  entry_id, 
  match_id,
  team_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetSelection :one
SELECT * FROM selections
WHERE id = $1 LIMIT 1;

-- name: ListSelections :many
SELECT * FROM selections
WHERE entry_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;