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

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	UserCreate := &models.User{}
	utils.ParseBody(r, UserCreate)
	u := UserCreate.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	NewUser := models.GetAllUser()
	res, _ := json.Marshal(NewUser)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	users := mux.Vars(r)
	userId := users["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing GetUserById")
	}
	userDetails, _ := models.GetUserById(Id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUpdate := models.User{}
	utils.ParseBody(r, userUpdate)
	users := mux.Vars(r)
	userId := users["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing UpdateUser")
	}
	userDetails, db := models.GetUserById(Id)
	if userUpdate.Name != "" {
		userDetails.Name = userUpdate.Name
	}
	if userUpdate.Username != "" {
		userDetails.Username = userUpdate.Username
	}
	if userUpdate.Email != "" {
		userDetails.Email = userUpdate.Email
	}
	if userUpdate.Password != "" {
		userDetails.Password = userUpdate.Password
	}

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	users := mux.Vars(r)
	userId := users["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing DeleteUser")
	}
	user := models.DeleteUserId(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}