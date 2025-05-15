-- name: GetSelection :one
SELECT * FROM selections
WHERE id = $1;

-- name: CreateSelection :one
INSERT INTO selections (
    entry_id, match_id, team_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: ChangeSelection :one
UPDATE selections 
SET
    match_id = $2,
    team_id = $3
WHERE id = $1
RETURNING *;

-- name: UpdateSelection :one
UPDATE selections
SET 
    is_correct = $2
WHERE id = $1
RETURNING *;