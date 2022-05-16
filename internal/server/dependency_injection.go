package server

import (
	"net/http"

	"github.com/parwalayush85/hands_on_go/internal/handler"
	"github.com/parwalayush85/hands_on_go/internal/service"
)

type userServiceApplication struct {
	userserverHttpHandler http.Handler
}

func newUserServerApplication() (*userServiceApplication, error) {
	userController := handler.NewUserController(
		&handler.UserValidatorImpl{},
		&service.UserServiceImpl{},
	)
	userserverHttpHandler := handler.UserServiceHttpHandler(userController)

	return &userServiceApplication{
		userserverHttpHandler: userserverHttpHandler,
	}, nil
}
