package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	Id, err :=strconv.ParseInt(hotelId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing GetHotelById")
	}
	hotelDetails,_ := models.GetHotelById(Id)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	hotelUpdate := models.Hotel{}
	utils.ParseBody(r, hotelUpdate)
	hotel := mux.Vars(r)
	hotelId := hotel["hotelId"]
	Id, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil{
		fmt.Println("Error while parsing UpdateHotel")
	}

	hotelDetails, db := models.GetHotelById(Id)
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
	if hotelUpdate.Title != "" {
		hotelDetails.Title = hotelUpdate.Title
	}

	db.Save(&hotelDetails)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	hotel := mux.Vars(r)
	hotelId := hotel["hotelid"]
	Id, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing DeleteHotel")
	}

	hotels := models.DeleteHotelId(Id)
	res, _ := json.Marshal(hotels)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}