package handler

import (
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
	switch blerr.GetKind(err) {
	case blerr.KindNotFound:
		return 404, userMsgOrDefault(err, "User Not Found")
	case blerr.KindInvalidInput:
		return 400, userMsgOrDefault(err, "Invalid input")
	default:
		return 500, userMsgOrDefault(err, "Internal Server Error")
	}
}
func userMsgOrDefault(err error, defaultMsg string) string {
	if userMsg, ok := blerr.GetUserMessageError(err); ok {
		return userMsg
	}
	return defaultMsg
}
