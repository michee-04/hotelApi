package controller

import (
	"encoding/json"
	"net/http"

	"github.com/michee/pkg/models"
	"github.com/michee/pkg/utils"
)

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	hotelCreate := &models.Hotel{}
	utils.ParseBody(r, hotelCreate)
	h := hotelCreate.CreateHotel()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetHotel(w http.ResponseWriter, r *http.Request) {

}
func GetHotelById(w http.ResponseWriter, r *http.Request) {

}
func UpdateHotel(w http.ResponseWriter, r *http.Request) {

}
func DeleteHotel(w http.ResponseWriter, r *http.Request) {

}