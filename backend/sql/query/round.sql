-- name: GetRound :one
SELECT * FROM rounds
WHERE id = $1;
