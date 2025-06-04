package user

import (
	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

func convertDBUsersToModelsUsers(users []db.User) *models.Users {
	mUsers := make(models.Users, len(users))
	for i := range users {
		mUsers[i] = convertDBUserToModelsUser(users[i])
	}

	return &mUsers
}

func convertDBUserToModelsUser(user db.User) *models.User {
	var phoneNumber string
	if user.PhoneNumber.Valid {
		phoneNumber = user.PhoneNumber.String
	}

	return &models.User{
		ID:            user.ID,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		PhoneNumber:   phoneNumber,
		DateOfBirth:   user.DateOfBirth,
		FavouriteTeam: &user.FavouriteTeamID.String,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
