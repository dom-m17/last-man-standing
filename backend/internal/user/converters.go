package user

import (
	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/models"
)

func convertDBUserToModelsUser(user db.User) *models.User {
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
	}
}
