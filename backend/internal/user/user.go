package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

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
	//TODO: Validation, casing, etc (ie any logic needed before inserting to DB)
	dob, _ := time.Parse("2006-01-02", input.DateOfBirth)

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		return &models.User{}, fmt.Errorf("hashing password: %w", err)
	}

	user, err := s.Querier.CreateUser(ctx, db.CreateUserParams{
		Username:       input.Username,
		HashedPassword: string(hashedPassword),
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

// LMS-45
func (s *Service) Login(ctx context.Context, input graphmodels.LoginInput) (*models.User, string, string, error) {
	user, err := s.Querier.GetUserByUsername(ctx, input.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.User{}, "", "", fmt.Errorf("no user exists with that username")
		}
		return &models.User{}, "", "", fmt.Errorf("getting user by username: %w", err)
	}

	if !checkPassword(input.Password, user.HashedPassword) {
		return &models.User{}, "", "", errors.New("password is incorrect")
	}
	// create refresh token
	s.Querier.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		UserID:    user.ID,
		TokenHash: s.RefreshToken.CreateRefreshToken(),
	})

	// create access token
	// use golang jwt?
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// return access token, refresh token and success
	return nil, token.Raw, "", nil
}
