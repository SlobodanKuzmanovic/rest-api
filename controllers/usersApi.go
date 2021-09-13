package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/models"
)

func UsersIndex(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	Users, err := models.AllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Users)
}
func HotUsersApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	Users, err := models.HotUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Users)
}

func LoginUser(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	var us models.LogUser
	_ = json.NewDecoder(r.Body).Decode(&us)
	fmt.Println(us)

	json.NewEncoder(w).Encode(models.OneUser(us.Email,us.Password))
}


func GetUserById(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)
	id:= vars["id"]

	json.NewEncoder(w).Encode(models.OneUserById(id))
}


func UpdateUserApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	var us models.User
	_ = json.NewDecoder(r.Body).Decode(&us)

	models.UpdateUser(us)

	json.NewEncoder(w).Encode(us)
}
func UpdateUserPasApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	var us models.User
	_ = json.NewDecoder(r.Body).Decode(&us)

	models.UpdateUserPass(us)

	json.NewEncoder(w).Encode(us)
}
func DeleteFromUsers(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)

	id:= vars["id"]

	models.DeleteUser(id)
}
func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/json")

	var us models.User
	_ = json.NewDecoder(r.Body).Decode(&us)

	models.NewUser(us)

	json.NewEncoder(w).Encode(us)
}

