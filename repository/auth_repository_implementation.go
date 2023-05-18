package repository

import (
	"nom/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DbRepo struct {
	DB *gorm.DB
}

func (db *DbRepo) FindUserById(id int) (*models.User, error) {
	var user models.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Return a 404 response if user is not found
			return nil, fiber.ErrNotFound
		}
		// Return an error if there's a database error
		return nil, result.Error
	}

	return &user, nil
}
