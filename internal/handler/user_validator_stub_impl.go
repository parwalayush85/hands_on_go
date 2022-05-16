package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/parwalayush85/hands_on_go/internal/blerr"
	"github.com/sirupsen/logrus"
)

type UserValidatorImpl struct {
}

func (u *UserValidatorImpl) ValidateGetUserById(r *http.Request) (int, error) {
	id, err := validateIdPaths(nil, r)
	if err != nil {
		return 0, blerr.ErrInvalidInput
	}
	return id, nil
}
func validateIdPaths(err error, r *http.Request) (int, error) {
	if err != nil {
		return 0, err
	}
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("id is required")
	}
	val, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("check id sent")
	}
	return val, nil
}
func (u *UserValidatorImpl) ValidateNewUser(r *http.Request) (*CreateUserRequest, error) {
	var user CreateUserRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&user)
	// logrus.Info(user)
	err = validateInputs(err, user.Age, "Age")
	err = validateInputs(err, user.FirstName, "FirstName")
	err = validateInputs(err, user.IsPhoneVerified, "IsPhoneVerified")
	err = validateInputs(err, user.LastName, "LastName")
	err = validateInputs(err, user.PhoneNumber, "PhoneNumber")
	if err != nil {
		return nil, blerr.ErrInvalidInput
	}
	if len(*user.FirstName) <= 0 || len(*user.FirstName) > 100 {
		return nil, blerr.ErrInvalidInput
	}
	if len(*user.LastName) <= 0 || len(*user.LastName) > 100 {
		return nil, blerr.ErrInvalidInput
	}
	logrus.Info(len(*user.PhoneNumber))
	if len(*user.PhoneNumber) <= 0 || len(*user.PhoneNumber) > 25 {
		return nil, blerr.ErrInvalidInput
	}
	// if reflect.TypeOf(user.Age) ==int{

	// }

	// err = validateInputSize(err, user.FirstName, "FirstName", 100, 0)
	return &user, nil
}

func validateInputs[T any](err error, value *T, valueName string) error {
	if err != nil {
		return err
	}
	if value == nil {
		return fmt.Errorf("%s is required", valueName)
	}
	return nil
}

// func validateInputSize[T any](err error, value *T, valueName string, maxSize int, minSize int) error {
// 	if err != nil {
// 		return err
// 	}

// 	if len(value)>=minSize&& len(value)<=maxSize{

// 	}
// 	return nil
// }
