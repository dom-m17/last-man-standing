-- name: GetRefreshTokenByTokenHash :one
SELECT * FROM refresh_tokens
WHERE token_hash = $1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (user_id, token_hash)
VALUES ($1, $2)
RETURNING *;