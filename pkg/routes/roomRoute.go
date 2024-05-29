package routes

import (
	"github.com/gorilla/mux"
	"github.com/michee/pkg/controller"
)


func RegisterRoom(route *mux.Router){

	route.HandleFunc("/room/", controller.CreateRoom).Methods("POST")
	route.HandleFunc("/room/", controller.GetRoom).Methods("GET")
	route.HandleFunc("/room/{roomId}", controller.GetRoomById).Methods("GET")
	route.HandleFunc("/room/{roomId}", controller.UpdateRoom).Methods("PUT")
	route.HandleFunc("/room/{roomId}", controller.DeleteRoom).Methods("DELETE")

}