package server

import (
	"net/http"

	"github.com/parwalayush85/hands_on_go/internal/handler"
)

type userServiceApplication struct {
	userserverHttpHandler http.Handler
}

func newUserServerApplication() (*userServiceApplication, error) {
	userserverHttpHandler := handler.UserServiceHttpHandler()

	return &userServiceApplication{
		userserverHttpHandler: userserverHttpHandler,
	}, nil
}
