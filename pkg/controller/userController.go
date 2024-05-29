package controller

import (
	"encoding/json"
	"net/http"

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

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}