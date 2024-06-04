package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michee/pkg/models"
	"github.com/michee/pkg/utils"
	"golang.org/x/crypto/bcrypt"
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
	userDetails, _ := models.GetUserById(userId)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUpdate := models.User{}
	utils.ParseBody(r, &userUpdate)
	users := mux.Vars(r)
	userId := users["userId"]
	userDetails, db := models.GetUserById(userId)
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
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.DefaultCost)
		userDetails.Password = string(hashedPassword)
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
	user := models.DeleteUserId(userId)
	res, _ := json.Marshal(user)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
