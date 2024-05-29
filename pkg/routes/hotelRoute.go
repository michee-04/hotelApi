package routes

import (
	"github.com/gorilla/mux"
	"github.com/michee/pkg/controller"
)


func RegisterHotel(route *mux.Router) {

	route.HandleFunc("/hotel", controller.CreateHotel).Methods("POST")
	route.HandleFunc("/hotel", controller.GetHotel).Methods("GET")
	route.HandleFunc("/hotel/{hotelId}", controller.GetHotelById).Methods("GET")
	route.HandleFunc("/hotel/{hotelId}", controller.UpdateHotel).Methods("PUT")
	route.HandleFunc("/hotel/{hotelId}", controller.DeleteHotel).Methods("DELETE")

}