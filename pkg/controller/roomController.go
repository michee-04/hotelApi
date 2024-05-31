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
	Id, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	roomDetails, _ := models.GetRoomById(Id)
	res, _ := json.Marshal(roomDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	roomUpdate := models.Room{}
	utils.ParseBody(r, roomUpdate)
	room := mux.Vars(r)
	roomId := room["roomId"]
	Id, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	roomDetails, db := models.GetRoomById(Id)
	if roomUpdate.Title != "" {
		roomDetails.Title = roomUpdate.Title
	}
	if roomUpdate.Image != "" {
		roomDetails.Image = roomUpdate.Image
	}
	if roomUpdate.Description != "" {
		roomDetails.Description = roomUpdate.Description
	}
	if roomUpdate.Title != "" {
		roomDetails.Title = roomUpdate.Title
	}
	if roomUpdate.RoomPrice != "" {
		roomDetails.RoomPrice = roomUpdate.RoomPrice
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
	Id, err := strconv.ParseInt(roomId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	rooms := models.DeleteRoomById(Id)
	res, _ := json.Marshal(rooms)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
