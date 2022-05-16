package handler

import (
	"net/http"

	"github.com/parwalayush85/hands_on_go/internal/blerr"
)

type UserValidatorImpl struct {
}

func (u *UserValidatorImpl) ValidateGetUserById(r *http.Request) (int, error) {
	if r.URL.Query().Has("invalid") {
		return 0, blerr.ErrInvalidInput
	}
	return 1, nil
}
