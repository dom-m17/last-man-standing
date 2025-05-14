-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1;

-- name: CreateEntry :one
INSERT INTO entries (
    user_id, competition_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateEntry :one
UPDATE entries
SET 
    status = $2
WHERE id = $1
RETURNING *;