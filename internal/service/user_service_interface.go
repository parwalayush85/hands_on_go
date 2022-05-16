package service

import (
	"github.com/parwalayush85/hands_on_go/internal/models"
)

type UserService interface {
	GetUserDetailsById(id int) (*models.User, error)
	DeleteUserById(id int) error
	CreateNewUser(user *models.User) (int, error)
}
