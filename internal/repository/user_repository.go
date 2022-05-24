package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/parwalayush85/hands_on_go/internal/models"
)

type UserRepository interface {
	checkExistsById(userId int) (bool, error)
	deleteById(userId int) error
	getUserById(userId int) (*models.User, error)
	CreateNewUser(user *models.User) (int, error)
}
type UserRepositoryImpl struct {
	db *sqlx.DB
}

func newUserRepository(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
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
func (u *UserRepositoryImpl) CreateNewUser(user *models.User) (int, error) {
	const query = `INSERT INTO USERS(first_name,last_name,age,phone_number,phone_verification_status) VALUES(?,?,?,?,?)`
	res, err := u.db.Exec(query, user.FirstName, user.LastName, user.Age, user.PhoneNumber, user.IsPhoneVerified)
	if err != nil {
		return 0, fmt.Errorf("DB Insert error")
	}
	newUserId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Last Insert Id error")
	}
	return int(newUserId), nil
}
