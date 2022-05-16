package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/parwalayush85/hands_on_go/internal/blerr"
	"github.com/parwalayush85/hands_on_go/internal/models"
	"github.com/parwalayush85/hands_on_go/internal/service"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	validator   UserValidator //controller depends on validator
	userService service.UserService
}

func NewUserController(uservalidator UserValidator, userService service.UserService) *UserController {
	return &UserController{
		validator:   uservalidator,
		userService: userService,
	}
}

func (u *UserController) GetUserById(writer http.ResponseWriter, reader *http.Request) {
	id, err := u.validator.ValidateGetUserById(reader)
	if errors.Is(err, blerr.ErrInvalidInput) {
		writer.WriteHeader(400)
		writer.Write([]byte("Request not valid"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error while validating user")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}
	user, err := u.userService.GetUserDetailsById(id)
	if errors.Is(err, blerr.ErrUserNotFound) {
		writer.WriteHeader(404)
		writer.Write([]byte("User Not Found"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error while getting user Info")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}
	getUserByIdResponse := u.toGetByIDResponse(user)
	getUserById, err := json.Marshal(getUserByIdResponse)
	if err != nil {
		logrus.WithError(err).Error("error while getting mapping user Info")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}
	writer.Write(getUserById)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	logrus.Info(user)
}

func (u *UserController) DeleteUserById(writer http.ResponseWriter, reader *http.Request) {
	id, err := u.validator.ValidateGetUserById(reader)
	if errors.Is(err, blerr.ErrInvalidInput) {
		writer.WriteHeader(400)
		writer.Write([]byte("Request not valid"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error while validating user")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}
	err = u.userService.DeleteUserById(id)
	if errors.Is(err, blerr.ErrUserNotFound) {
		writer.WriteHeader(404)
		writer.Write([]byte("User Not Found"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error while getting user Info")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}
func (u *UserController) CreateNewUser(writer http.ResponseWriter, reader *http.Request) {
	id, err := u.validator.ValidateNewUser(reader)
	if errors.Is(err, blerr.ErrInvalidInput) {
		writer.WriteHeader(400)
		writer.Write([]byte("Request not valid"))
		return
	}
	if err != nil {
		logrus.WithError(err).Error("error while validating user")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}

	userId, err := u.userService.CreateNewUser(toGetByIDResponse(id))
	if err != nil {
		logrus.WithError(err).Error("error while updating user")
		writer.WriteHeader(500)
		writer.Write([]byte("An unexpected error has occured"))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(strconv.Itoa(userId)))
}

func (u *UserController) toGetByIDResponse(user *models.User) *getUserByIdResponse {
	return &getUserByIdResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Age:             user.Age,
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.IsPhoneVerified,
	}
}
func toGetByIDResponse(user *CreateUserRequest) *models.User {
	return &models.User{
		ID:              123,
		Age:             jsonNumbertoInt(*user.Age),
		FirstName:       *user.FirstName,
		LastName:        *user.LastName,
		PhoneNumber:     *user.PhoneNumber,
		IsPhoneVerified: *user.IsPhoneVerified,
	}
}
func jsonNumbertoInt(n json.Number) int {
	id, _ := strconv.Atoi(n.String())
	return id
}
