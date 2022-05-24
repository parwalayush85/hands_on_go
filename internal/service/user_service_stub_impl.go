package service

import (
	"fmt"

	"github.com/parwalayush85/hands_on_go/internal/blerr"
	"github.com/parwalayush85/hands_on_go/internal/models"
)

type UserRepository interface {
	checkExistsById(userId int) (bool, error)
	deleteById(userId int) error
	getUserById(userId int) (*models.User, error)
	CreateNewUser(user *models.User) (int, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserServiceImpl(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u *UserServiceImpl) GetUserDetailsById(id int) (*models.User, error) {
	userExists, err := u.userRepository.checkExistsById(id)
	if err != nil {
		return nil, ErrorReturn(fmt.Errorf("something went wrong while checking user exists"), blerr.KindInternalServerError)
	}
	if !userExists {
		return nil, ErrorReturn(fmt.Errorf("user Not Found"), blerr.KindNotFound)
	}
	user, err := u.userRepository.getUserById(id)
	if err != nil {
		return nil, ErrorReturn(fmt.Errorf("something went wrong while getting user value"), blerr.KindInternalServerError)
	}
	// return &models.User{ID: 1, FirstName: "Ayush", LastName: "Parwal", Age: 23, PhoneNumber: "9804710111", IsPhoneVerified: true}, nil
	return user, nil
}
func (u *UserServiceImpl) DeleteUserById(id int) error {
	userExists, err := u.userRepository.checkExistsById(id)
	if err != nil {
		return ErrorReturn(fmt.Errorf("something went wrong while checking user exists"), blerr.KindInternalServerError)
	}
	if !userExists {
		return ErrorReturn(fmt.Errorf("user Not Found"), blerr.KindNotFound)
	}
	err = u.userRepository.deleteById(id)
	if err != nil {
		return ErrorReturn(fmt.Errorf("something went wrong while deleting user"), blerr.KindInternalServerError)
	}

	return nil
}
func (u *UserServiceImpl) CreateNewUser(user *models.User) (int, error) {
	u.userRepository.CreateNewUser(user)
	return user.ID, nil
}
func ErrorReturn(err error, kind blerr.Kind) error {
	err = blerr.SetUserMsgError(err, err.Error())
	return blerr.SetKind(err, kind)
}
