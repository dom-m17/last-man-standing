-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, 
  hashed_password, 
  first_name, 
  last_name, 
  email, 
  phone_number, 
  date_of_birth,
  favourite_team
) VALUES (
 $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  hashed_password = $3,
  first_name = $4,
  last_name = $5,
  email = $6,
  phone_number = $7,
  date_of_birth = $8,
  favourite_team = $9
WHERE id = $1
RETURNING *;