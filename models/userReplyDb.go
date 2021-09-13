package models

import "fmt"

type UserReply struct {
	Pk_UserReplyId string `json:"pk_UserReplyId"`
	Reaction bool `json:"reaction"`
	Fk_UserId string `json:"fk_UserId"`
	Fk_ReplyId string `json:"fk_ReplyId"`
}

func NewReactionOnReply(react UserReply)  {
	results, err := DB.Query("SELECT Pk_UserReplyId FROM `usersreplies` WHERE Fk_UserId = ? AND Fk_ReplyId =?", react.Fk_UserId, react.Fk_ReplyId)
	if err != nil {
		panic(err.Error())
	}
	var reaction *UserReply
	for results.Next() {
		reaction = new(UserReply)
		err = results.Scan(&reaction.Pk_UserReplyId)
		if err != nil {
			panic(err.Error())
		}
	}
	if reaction == nil	{
		query := fmt.Sprintf("INSERT INTO `usersreplies`(`Reaction`, `Fk_UserId`, `Fk_ReplyId`) VALUES (%t, '%s','%s')",react.Reaction,react.Fk_UserId,react.Fk_ReplyId)
		_, err := DB.Query(query)
		if err != nil {
			panic(err.Error())
		}
	} else {
		query := fmt.Sprintf("UPDATE `usersreplies` SET `Reaction`= %t WHERE Pk_UserReplyId = %s", react.Reaction, reaction.Pk_UserReplyId)

		_, err := DB.Query(query)
		if err != nil {
			panic(err.Error())
		}
	}

}
