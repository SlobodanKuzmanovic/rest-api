package models

import "fmt"

type UserQuestion struct {
	Pk_UserQuestionId string `json:"pk_UserQuestionId"`
	Reaction bool `json:"reaction"`
	Fk_UserId string `json:"fk_UserId"`
	Fk_QuestionId string `json:"fk_QuestionId"`
}

func InsertReaction(reac UserQuestion) {
	results, err := DB.Query("SELECT Pk_UserQuestionId FROM `usersquestions` WHERE Fk_UserId = ? AND Fk_QuestionId =?", reac.Fk_UserId, reac.Fk_QuestionId)
	if err != nil {
		panic(err.Error())
	}
	var reaction *UserQuestion
	for results.Next() {
		reaction = new(UserQuestion)
		err = results.Scan(&reaction.Pk_UserQuestionId)
		if err != nil {
			panic(err.Error())
		}
	}
	if reaction == nil	{
		query := fmt.Sprintf("INSERT INTO `usersquestions`(`Reaction`, `Fk_UserId`, `Fk_QuestionId`) VALUES (%t, '%s','%s')", reac.Reaction, reac.Fk_UserId, reac.Fk_QuestionId)

		_, err := DB.Query(query)
		if err != nil {
			panic(err.Error())
		}
	} else {
		query := fmt.Sprintf("UPDATE `usersquestions` SET `Reaction`= %t WHERE Pk_UserQuestionId = %s", reac.Reaction, reaction.Pk_UserQuestionId)

		_, err := DB.Query(query)
		if err != nil {
			panic(err.Error())
		}
	}
}