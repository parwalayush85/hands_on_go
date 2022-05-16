package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserServiceHttpHandler(u *UserController) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", u.GetUserById).Methods("GET")
	return r
	// return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	// 	logrus.Info("Here")
	// 	writer.Write([]byte("Hello World"))
	// 	// fmt.Printf("Hello")
	// })
}
