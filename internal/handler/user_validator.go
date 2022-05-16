package handler

import "net/http"

type UserValidator interface {
	ValidateGetUserById(r *http.Request) (int, error)
}
