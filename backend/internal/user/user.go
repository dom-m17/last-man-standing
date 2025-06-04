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

func (s *Service) ListUsers(ctx context.Context) (models.Users, error) {
	users, err := s.Querier.ListUsers(ctx)
	if err != nil {
		return models.Users{}, fmt.Errorf("listing users: %w", err)
	}

	return *convertDBUsersToModelsUsers(users), nil
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
		PhoneNumber: sql.NullString{
			String: func() string {
				if input.PhoneNumber != nil {
					return *input.PhoneNumber
				}
				return ""
			}(),
			Valid: input.PhoneNumber != nil && *input.PhoneNumber != "",
		},
		DateOfBirth: dob,
		FavouriteTeamID: sql.NullString{
			String: func() string {
				if input.FavouriteTeam != nil {
					return *input.FavouriteTeam
				}
				return ""
			}(),
			Valid: input.FavouriteTeam != nil && *input.FavouriteTeam != "",
		},
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

func (s *Service) UpdateUser(ctx context.Context, input graphmodels.UpdateUserInput) (*models.User, error) {
	dob, err := time.Parse("2006-01-02", input.DateOfBirth)
	if err != nil {
		return &models.User{}, fmt.Errorf("invalid date of birth format: %w", err)
	}
	user, err := s.Querier.UpdateUser(ctx, db.UpdateUserParams{
		ID:        input.ID,
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		PhoneNumber: sql.NullString{
			String: func() string {
				if input.PhoneNumber != nil {
					return *input.PhoneNumber
				}
				return ""
			}(),
			Valid: input.PhoneNumber != nil && *input.PhoneNumber != "",
		},
		DateOfBirth: dob,
		FavouriteTeamID: sql.NullString{
			String: func() string {
				if input.FavouriteTeam != nil {
					return *input.FavouriteTeam
				}
				return ""
			}(),
			Valid: input.FavouriteTeam != nil && *input.FavouriteTeam != "",
		},
	})
	if err != nil {
		return &models.User{}, fmt.Errorf("deleting user: %w", err)
	}

	return convertDBUserToModelsUser(user), nil
}
