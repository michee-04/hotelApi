package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michee/pkg/models"
	"github.com/michee/pkg/utils"
)

var NewHotel models.Hotel

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	hotelCreate := &models.Hotel{}
	utils.ParseBody(r, hotelCreate)
	h := hotelCreate.CreateHotel()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetHotel(w http.ResponseWriter, r *http.Request) {
	NewHotel := models.GetAllHotel()
	res, _ := json.Marshal(NewHotel)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetHotelById(w http.ResponseWriter, r *http.Request) {
	hotel := mux.Vars(r)
	hotelId := hotel["hotelId"]
	hotelDetails,_ := models.GetHotelById(hotelId)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	hotel := mux.Vars(r)
	hotelId := hotel["hotelId"]
	hotelUpdate := &models.Hotel{}
	utils.ParseBody(r, hotelUpdate)

	hotelDetails, db := models.GetHotelById(hotelId)

	if hotelUpdate.Title != "" {
			hotelDetails.Title = hotelUpdate.Title
	}
	if hotelUpdate.Image != "" {
			hotelDetails.Image = hotelUpdate.Image
	}
	if hotelUpdate.Description != "" {
			hotelDetails.Description = hotelUpdate.Description
	}
	if hotelUpdate.Localisation != "" {
			hotelDetails.Localisation = hotelUpdate.Localisation
	}
	if hotelUpdate.City != "" {
			hotelDetails.City = hotelUpdate.City
	}
	if hotelUpdate.State != "" {
			hotelDetails.State = hotelUpdate.State
	}

	db.Save(&hotelDetails)

	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	hotel := mux.Vars(r)
	hotelId := hotel["hotelId"]
	deletedHotel := models.DeleteHotelId(hotelId)

	res, _ := json.Marshal(deletedHotel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}