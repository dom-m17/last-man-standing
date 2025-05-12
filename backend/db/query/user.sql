-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, 
  hashed_password, 
  first_name, 
  last_name, 
  email, 
  phone_number, 
  favourite_team
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;