package services

import (
	"nom/models"
	"nom/repository"
)

type AuthService struct {
	AuthRepo repository.AuthRepository
}

func (auth *AuthService) GetUserById(id int) (*models.User, error) {
	return auth.AuthRepo.FindUserById(id)
}

// func FindUserByEmail(email string, user *models.User) error {
// 	database.Database.DB.Find(&user, "email = ?", email)

// 	if user.ID == 0 {
// 		return errors.New("user does not exsts")
// 	}

// 	return nil
// }
