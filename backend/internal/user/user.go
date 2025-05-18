package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	"github.com/dom-m17/lms/backend/internal/subgraph/model"
)

func (s *Service) GetUser(ctx context.Context, input string) (*models.User, error) {
	user, err := s.Querier.GetUser(ctx, input)
	if err != nil {
		fmt.Printf("DB error: %+v\n", err)
		return &models.User{}, fmt.Errorf("getting user: %w", err)
	}

	return &models.User{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	//TODO: Validation, hashing, etc (ie any logic needed before inserting to DB)
	user, err := s.Querier.CreateUser(ctx, db.CreateUserParams{
		Username:       input.Username,
		HashedPassword: input.HashedPassword,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		PhoneNumber:    sql.NullString{String: input.PhoneNumber, Valid: input.PhoneNumber != ""},
	})
	if err != nil {
		return &models.User{}, fmt.Errorf("creating user: %w", err)
	}

	var phoneNumber string
	if user.PhoneNumber.Valid {
		phoneNumber = user.PhoneNumber.String
	}

	return &models.User{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: phoneNumber,
	}, nil
}
