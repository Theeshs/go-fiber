package repository

import (
	"nom/models"
)

type AuthRepository interface {
	FindUserById(id int) (*models.User, error)
}
