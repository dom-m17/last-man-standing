package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dom-m17/lms/backend/internal/utils"
	"github.com/peterldowns/testy/check"
)

// This is a helper function because any other test will require creating a user first
func createTestUser(t *testing.T, ctx context.Context, q Querier) User {
	userToCreate := CreateUserParams{
		Username:       utils.RandomUsername(),
		HashedPassword: utils.RandomPassword(),
		FirstName:      utils.RandomString(5),
		LastName:       utils.RandomString(5),
		Email:          utils.RandomEmail(),
		PhoneNumber:    sql.NullString{String: utils.RandomPhoneNumber(), Valid: true},
		DateOfBirth:    utils.RandomDateOfBirth(),
	}

	createdUser, err := q.CreateUser(ctx, userToCreate)
	check.Nil(t, err)
	check.Equal(t, createdUser.Username, userToCreate.Username)
	check.Equal(t, createdUser.HashedPassword, userToCreate.HashedPassword)
	check.Equal(t, createdUser.FirstName, userToCreate.FirstName)
	check.Equal(t, createdUser.LastName, userToCreate.LastName)
	check.Equal(t, createdUser.Email, userToCreate.Email)
	check.Equal(t, createdUser.PhoneNumber.String, userToCreate.PhoneNumber.String)
	check.Equal(t, createdUser.DateOfBirth, userToCreate.DateOfBirth)

	return createdUser
}

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	q, close := NewTestQuerier(t)
	defer close()
	createTestUser(t, ctx, q)
}

func Test_GetUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	q, close := NewTestQuerier(t)
	defer close()

	createdUser := createTestUser(t, ctx, q)
	gotUser, err := q.GetUser(ctx, createdUser.ID)
	check.Nil(t, err)
	check.Equal(t, createdUser, gotUser)
}

func Test_DeleteUser(t *testing.T) {
	t.Parallel()
	ctx := t.Context()

	q, close := NewTestQuerier(t)
	defer close()

	createdUser := createTestUser(t, ctx, q)
	deletedUser, err := q.DeleteUser(ctx, createdUser.ID)
	check.Nil(t, err)
	check.Equal(t, createdUser, deletedUser)
	_, err = q.GetUser(ctx, deletedUser.ID)
	check.Error(t, err)
}

func Test_ListUsers(t *testing.T) {
	// TODO
}

func Test_UpdateUsers(t *testing.T) {
	// TODO
}
