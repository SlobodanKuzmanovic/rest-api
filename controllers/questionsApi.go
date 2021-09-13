package controllers

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/models"
	"strconv"
)

func QuestionsIndex(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	Questions, err := models.AllQuestions()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Questions)
}
func HotQuestionsApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	Questions, err := models.HotQuestions()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Questions)
}
func GetQuestionById(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)
	id:= vars["id"]

	json.NewEncoder(w).Encode(models.QuestionById(id))
}
func QuestionsPagingApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["n"])

	Questions, err := models.QuestionPaging(n)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Questions)
}
func GetQuestionsByUserIdApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)

	vars := mux.Vars(r)

	id:= vars["id"]
	n, err := strconv.Atoi(vars["n"])

	Questions, err := models.GetQuestionsByUserId(id,n)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	json.NewEncoder(w).Encode(Questions)
}

func DeleteFromQuestions(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)

	vars := mux.Vars(r)

	id:= vars["id"]

	models.DeleteQuestion(id)
}
func CreateNewQuestion(w http.ResponseWriter, r *http.Request) {
	SetupCorsResponse(&w, r)

	w.Header().Set("Content-Type", "application/json")

	var quest models.Question
	_ = json.NewDecoder(r.Body).Decode(&quest)

	models.NewQuestion(quest)

	json.NewEncoder(w).Encode(quest)
}
func UpdateQuestionApi(w http.ResponseWriter, r *http.Request)  {
	SetupCorsResponse(&w, r)

	var quest models.Question
	_ = json.NewDecoder(r.Body).Decode(&quest)

	models.UpdateQuestion(quest)

	json.NewEncoder(w).Encode(quest)
}