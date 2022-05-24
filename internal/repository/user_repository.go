package repository

import (
	"errors"

	"github.com/parwalayush85/hands_on_go/internal/models"
)

type UserRepository interface {
	checkExistsById(userId int) (bool, error)
	deleteById(userId int) error
	getUserById(userId int) (*models.User, error)
}
type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) checkExistsById(userId int) (bool, error) {
	if userId == 404 {
		return false, nil
	}
	if userId == 500 {
		return false, errors.New("Internal server")
	}
	return true, nil
}
func (u *UserRepositoryImpl) deleteById(userId int) error {
	if userId == 123 {
		return nil
	}
	return errors.New("Internal server")
}
func (u *UserRepositoryImpl) getUserById(userId int) (*models.User, error) {
	if userId == 123 {
		return &models.User{ID: 123, FirstName: "Ayush", LastName: "Parwal", Age: 23, PhoneNumber: "9804710111", IsPhoneVerified: true}, nil
	}
	return nil, errors.New("Internal server")
}
