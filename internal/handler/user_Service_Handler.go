package handler

import (
	"github.com/gorilla/mux"

	"net/http"
)

func UserServiceHttpHandler(u *UserController) http.Handler {
	r := mux.NewRouter()
	r.Handle("/users/{id}", ErrResponseAdapter(u.GetUserById),).Methods("GET")
	r.Handle("/users/{id}", ErrResponseAdapter(u.DeleteUserById),).Methods("DELETE")
	r.HandleFunc("/users/new", u.CreateNewUser).Methods("POST")
	return r
	// return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	// 	logrus.Info("Here")
	// 	writer.Write([]byte("Hello World"))
	// 	// fmt.Printf("Hello")
	// })
}
