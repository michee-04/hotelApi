package routes

import (

	"github.com/gorilla/mux"
	"github.com/michee/pkg/controller"
)


func RegisterUser(route *mux.Router) {

	route.HandleFunc("/user", controller.CreateUser).Methods("POST")
	route.HandleFunc("/user", controller.GetUser).Methods("GET")
	route.HandleFunc("/user/{userId}", controller.GetUserById).Methods("GET")
	route.HandleFunc("/user/{userId}", controller.UpdateUser).Methods("PUT")
	route.HandleFunc("/user/{userId}", controller.DeleteUser).Methods("DELETE")

}