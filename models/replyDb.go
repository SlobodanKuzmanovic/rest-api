package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Reply struct {
	Pk_ReplyId string `json:"pk_ReplyId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
	Fk_QuestionId string `json:"fk_QuestionId"`
}
type ReplyWithReaction struct {
	Pk_ReplyId string `json:"pk_ReplyId"`
	Text string `json:"text"`
	Fk_UserId string `json:"fk_UserId"`
	Fk_QuestionId string `json:"fk_QuestionId"`
	Positive string `json:"positive"`
	Negative string `json:"negative"`
}
type Replies []Reply

func AllReplies ()([]Reply, error)  {

	results, err := DB.Query("SELECT Pk_ReplyId, Text, Fk_UserId, Fk_QuestionId from Replies")
	if err != nil {
		panic(err.Error())
	}

	var Replies []Reply

	for results.Next() {
		var rep Reply

		err = results.Scan(&rep.Pk_ReplyId, &rep.Text, &rep.Fk_UserId, &rep.Fk_QuestionId)
		if err != nil {
			panic(err.Error())
		}

		Replies = append(Replies, rep)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Replies, err
}
func RepliesByQuestionId(id string)([]ReplyWithReaction, error)  {

	results, err := DB.Query("SELECT r.Pk_ReplyId, r.Text, r.Fk_UserId, r.Fk_QuestionId, SUM(IF(u.Reaction = 1, 1, 0)) AS positive, " +
		"SUM(IF(u.Reaction = 0, 1, 0)) AS negative FROM replies as r JOIN usersreplies as u " +
		"ON u.Fk_ReplyId=r.Pk_ReplyId WHERE Fk_QuestionId = ?", id)
	if err != nil {
		panic(err.Error())
	}

	var Replies []ReplyWithReaction
	for results.Next() {
		var rep ReplyWithReaction
		err = results.Scan(&rep.Pk_ReplyId, &rep.Text, &rep.Fk_UserId, &rep.Fk_QuestionId, &rep.Positive, &rep.Negative)
		if err != nil {
			panic(err.Error())
		}
		Replies = append(Replies, rep)
		if err = results.Err(); err != nil {
			return nil, err
		}
	}
	return Replies, err
}
func ReplyById(id string)(ReplyWithReaction)  {

	results, err := DB.Query("SELECT r.Pk_ReplyId, r.Text, r.Fk_UserId, r.Fk_QuestionId, SUM(IF(u.Reaction = 1, 1, 0)) AS positive, " +
		"SUM(IF(u.Reaction = 0, 1, 0)) AS negative FROM replies as r JOIN usersreplies as u " +
		"ON u.Fk_ReplyId=r.Pk_ReplyId WHERE r.Pk_ReplyId = ?", id)
	if err != nil {
		panic(err.Error())
	}

	var rep ReplyWithReaction
	for results.Next() {

		err = results.Scan(&rep.Pk_ReplyId, &rep.Text, &rep.Fk_UserId, &rep.Fk_QuestionId, &rep.Positive, &rep.Negative)
		if err != nil {
			panic(err.Error())
		}
	}
	return rep
}
func DeleteReply(id string)()  {
	_, err := DB.Query("DELETE FROM `Replies` WHERE Pk_ReplyId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewReply (rep Reply) () {

	query := fmt.Sprintf("INSERT INTO `replies`(`Text`, `Fk_UserId`,`Fk_QuestionId`) VALUES ('%s', '%s', '%s')", rep.Text,rep.Fk_UserId, rep.Fk_QuestionId)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func UpdateReply (rep Reply) () {
	query := fmt.Sprintf("UPDATE `replies` SET `Text`= '%s', `Fk_UserId`= '%s',`Fk_QuestionId`='%s' WHERE Pk_ReplyId= %s", rep.Text, rep.Fk_UserId, rep.Fk_QuestionId, rep.Pk_ReplyId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
