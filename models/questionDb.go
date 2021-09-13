package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Question struct {
	Pk_QuestionId string `json:"pk_QuestionId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
}
type QuestionWithLikes struct {
	Pk_QuestionId string `json:"pk_QuestionId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
	Positive string `json:"positive"`
	Negative string `json:"negative"`
}
type QuestionHot struct {
	Pk_QuestionId string `json:"pk_QuestionId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
	ReactionCount string `json:"reaction_count"`
}

type Questions []Question
var paging = 20
func AllQuestions ()([]Question, error)  {

	results, err := DB.Query("SELECT Pk_QuestionId, Text, Fk_UserId from Questions")
	if err != nil {
		panic(err.Error())
	}

	var Questions []Question

	for results.Next() {
		var quest Question

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId)
		if err != nil {
			panic(err.Error())
		}

		Questions = append(Questions, quest)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Questions, err
}
func HotQuestions ()([]QuestionHot, error)  {

	results, err := DB.Query("SELECT q.Pk_QuestionId, q.Text, q.Fk_UserId, COUNT(l.Pk_UserQuestionId) as Hotest " +
		"FROM questions AS q JOIN usersquestions AS l ON q.Pk_QuestionId = l.Fk_QuestionId " +
		"GROUP BY q.Pk_QuestionId, q.Text, q.Fk_UserId ORDER BY hotest DESC LIMIT 5 OFFSET 0")
	if err != nil {
		panic(err.Error())
	}

	var Questions []QuestionHot

	for results.Next() {
		var quest QuestionHot

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId, &quest.ReactionCount)
		if err != nil {
			panic(err.Error())
		}

		Questions = append(Questions, quest)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Questions, err
}
func QuestionById(id string)(Question)  {

	results, err := DB.Query("SELECT Pk_QuestionId, Text, Fk_UserId FROM `questions` WHERE Pk_QuestionId = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var quest Question
	for results.Next() {

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	return quest
}
func QuestionPaging (n int)([]QuestionWithLikes, error)  {

	results, err := DB.Query("SELECT q.Pk_QuestionId, q.Text, q.Fk_UserId, SUM(IF(up.Reaction = 1, 1, 0)) AS positive, " +
		"SUM(IF(up.Reaction = 0, 1, 0)) AS negative FROM Questions AS q " +
		"LEFT JOIN usersquestions AS up ON q.Pk_QuestionId = up.Fk_QuestionId " +
		"GROUP BY q.Pk_QuestionId, q.Text, q.Fk_UserId " +
		"ORDER BY Pk_QuestionId DESC LIMIT ? OFFSET ?", paging,(n-1)*5)
	if err != nil {
		panic(err.Error())
	}

	var Questions []QuestionWithLikes

	for results.Next()	{
		var quest QuestionWithLikes

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId, &quest.Positive, &quest.Negative)
		if err != nil {
			panic(err.Error())
		}

		Questions = append(Questions, quest)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Questions, err
}

func GetQuestionsByUserId (id string, n int)([]Question, error)  {

	results, err := DB.Query("SELECT Pk_QuestionId, Text, Fk_UserId from Questions WHERE Fk_UserId = ? ORDER BY Pk_QuestionId DESC LIMIT ? OFFSET ?", id,paging, (n-1)*5)
	if err != nil {
		panic(err.Error())
	}

	var Questions []Question

	for results.Next()	{
		var quest Question

		err = results.Scan(&quest.Pk_QuestionId, &quest.Text, &quest.Fk_UserId)
		if err != nil {
			panic(err.Error())
		}

		Questions = append(Questions, quest)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Questions, err
}
func DeleteQuestion(id string)()  {
	_, err := DB.Query("DELETE FROM `Questions` WHERE Pk_QuestionId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewQuestion (quest Question) () {

	query := fmt.Sprintf("INSERT INTO `questions`(`Text`, `Fk_UserId`) VALUES ('%s', '%s')", quest.Text,quest.Fk_UserId)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func UpdateQuestion (quest Question) () {
	query := fmt.Sprintf("UPDATE `Questions` SET `Text`= '%s', `Fk_UserId`= '%s' WHERE Pk_QuestionId= %s", quest.Text, quest.Fk_UserId, quest.Pk_QuestionId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}