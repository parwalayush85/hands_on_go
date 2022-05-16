package service

import "github.com/parwalayush85/hands_on_go/internal/models"

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUserDetailsById(id int) (*models.User, error) {
	return &models.User{ID: 1, FirstName: "Ayush", LastName: "Parwal", Age: 23, PhoneNumber: 9804710111, IsPhoneVerified: true}, nil
}
