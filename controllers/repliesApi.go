package controllers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/models"
)

func RepliesIndex(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	Replies, err := models.AllReplies()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Replies)
}
func GetRepliesByQuestionId(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)
	id:= vars["id"]
	Replies, err := models.RepliesByQuestionId(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Replies)
}
func GetReplyByIdApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)

	id:= vars["id"]

	json.NewEncoder(w).Encode(models.ReplyById(id))
}
func DeleteFromReplies(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)

	id:= vars["id"]

	models.DeleteReply(id)
}
func CreateNewReply(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/json")

	var rep models.Reply
	_ = json.NewDecoder(r.Body).Decode(&rep)

	models.NewReply(rep)

	json.NewEncoder(w).Encode(rep)
}
func UpdateReplyApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)

	var rep models.Reply
	_ = json.NewDecoder(r.Body).Decode(&rep)

	models.UpdateReply(rep)

	json.NewEncoder(w).Encode(rep)
}
