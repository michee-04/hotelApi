package routes

import (
	"github.com/gorilla/mux"
	"github.com/michee/pkg/controller"
)

func RegisterBookings(route *mux.Router){

	route.HandleFunc("/booking", controller.CreateBooking).Methods("POST")

	route.HandleFunc("/booking", controller.GetBooking).Methods("GET")

	route.HandleFunc("/booking/{bookingId}", controller.GetBookingById).Methods("GET")
	
	route.HandleFunc("/booking/{bookingId}", controller.UpdateBooking).Methods("PUT")

	route.HandleFunc("/booking/{bookingId}", controller.DeleteBooking).Methods("DELETE")
}