// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
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
RETURNING id, username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team, created_at, updated_at
`

type CreateUserParams struct {
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	PhoneNumber    sql.NullString `json:"phone_number"`
	DateOfBirth    time.Time      `json:"date_of_birth"`
	FavouriteTeam  sql.NullInt64  `json:"favourite_team"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		arg.DateOfBirth,
		arg.FavouriteTeam,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.FavouriteTeam,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team, created_at, updated_at
`

func (q *Queries) DeleteUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.FavouriteTeam,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team, created_at, updated_at FROM users
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.FavouriteTeam,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team, created_at, updated_at FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.PhoneNumber,
			&i.DateOfBirth,
			&i.FavouriteTeam,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
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
RETURNING id, username, hashed_password, first_name, last_name, email, phone_number, date_of_birth, favourite_team, created_at, updated_at
`

type UpdateUserParams struct {
	ID             string         `json:"id"`
	Username       string         `json:"username"`
	HashedPassword string         `json:"hashed_password"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	PhoneNumber    sql.NullString `json:"phone_number"`
	DateOfBirth    time.Time      `json:"date_of_birth"`
	FavouriteTeam  sql.NullInt64  `json:"favourite_team"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		arg.DateOfBirth,
		arg.FavouriteTeam,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.DateOfBirth,
		&i.FavouriteTeam,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
