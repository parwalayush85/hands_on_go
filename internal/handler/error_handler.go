package handler

import (
	"errors"
	"net/http"

	"github.com/parwalayush85/hands_on_go/internal/blerr"
)

type ErrHttpHandler func(w http.ResponseWriter, r *http.Request) error

func ErrResponseAdapter(next ErrHttpHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		status, message := getStatusAndMsg(err)
		w.WriteHeader(status)
		w.Write([]byte(message))
	})
}
func getStatusAndMsg(err error) (int, string) {
	if errors.Is(err, blerr.ErrInvalidInput) {
		return 400, "request not valid"
	}
	if errors.Is(err, blerr.ErrUserNotFound) {
		return 404, "user not found"
	}
	return 500, "internal server error"
}
