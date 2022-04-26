package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func UserServiceHttpHandler() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logrus.Info("Here")
		writer.Write([]byte("Hello World"))
		// fmt.Printf("Hello")
	})
}
