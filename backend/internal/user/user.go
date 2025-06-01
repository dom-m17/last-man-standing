package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
	graphmodels "github.com/dom-m17/lms/backend/internal/subgraph/model"
)

func (s *Service) GetUser(ctx context.Context, input string) (*models.User, error) {
	user, err := s.Querier.GetUser(ctx, input)
	if err != nil {
		fmt.Printf("DB error: %+v\n", err)
		return &models.User{}, fmt.Errorf("getting user: %w", err)
	}

	return convertDBUserToModelsUser(user), nil
}

func (s *Service) CreateUser(ctx context.Context, input graphmodels.CreateUserInput) (*models.User, error) {
	//TODO: Validation, hashing, casing, etc (ie any logic needed before inserting to DB)
	dob, _ := time.Parse("2006-01-02", input.DateOfBirth)
	user, err := s.Querier.CreateUser(ctx, db.CreateUserParams{
		Username:       input.Username,
		HashedPassword: input.HashedPassword,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		PhoneNumber:    sql.NullString{String: input.PhoneNumber, Valid: input.PhoneNumber != ""},
		DateOfBirth:    dob,
	})
	if err != nil {
		return &models.User{}, fmt.Errorf("creating user: %w", err)
	}

	return convertDBUserToModelsUser(user), nil
}

func (s *Service) DeleteUser(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.Querier.DeleteUser(ctx, userID)
	if err != nil {
		return &models.User{}, fmt.Errorf("deleting user: %w", err)
	}

	return convertDBUserToModelsUser(user), nil
}
