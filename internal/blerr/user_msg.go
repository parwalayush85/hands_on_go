package blerr

import (
	"errors"
)

type UserMsgError struct {
	userMsg string
	err     error
}

func (ume *UserMsgError) Error() string {
	if ume.err == nil {
		return ""
	}
	return ume.err.Error()
}
func (ume *UserMsgError) Unwrap() error {
	return ume.err
}
func SetUserMsgError(err error, userMsg string) error {
	return &UserMsgError{userMsg: userMsg, err: err}
}
func GetUserMessageError(err error) (string, bool) {
	var userMsgError *UserMsgError
	if errors.As(err, &userMsgError) {
		return userMsgError.userMsg, true
	}
	return "Unknown Error Occured", false
}
func NewWithUserMsg(userMsg string) error {
	return &UserMsgError{userMsg: userMsg, err: errors.New(userMsg)}
}
