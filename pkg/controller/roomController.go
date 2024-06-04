package controller

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michee/pkg/models"
	"github.com/michee/pkg/utils"
)

var NewRoom models.Room

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	roomCreate := &models.Room{}
	utils.ParseBody(r, roomCreate)
	h := roomCreate.CreateRoom()
	res, _ := json.Marshal(h)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	rooms := models.GetAllRoom()
	res, _ := json.Marshal(rooms)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRoomById(w http.ResponseWriter, r *http.Request) {
	room := mux.Vars(r)
	roomId := room["roomId"]
	roomDetails, _ := models.GetRoomById(roomId)
	res, _ := json.Marshal(roomDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	roomUpdate := models.Room{}
	utils.ParseBody(r, &roomUpdate)
	room := mux.Vars(r)
	roomId := room["roomId"]
	roomDetails, db := models.GetRoomById(roomId)
	if roomUpdate.Title != "" {
		roomDetails.Title = roomUpdate.Title
	}
	if roomUpdate.Image != "" {
		roomDetails.Image = roomUpdate.Image
	}
	if roomUpdate.Description != "" {
		roomDetails.Description = roomUpdate.Description
	}
	if roomUpdate.RoomPrice != "" {
		roomDetails.RoomPrice = roomUpdate.RoomPrice
	}
	if roomUpdate.HotelID != "" {
		roomDetails.HotelID = roomUpdate.HotelID
	}
	db.Save(&roomDetails)
	res, _ := json.Marshal(roomDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	room := mux.Vars(r)
	roomId := room["roomId"]
	rooms := models.DeleteRoomById(roomId)
	res, _ := json.Marshal(rooms)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
