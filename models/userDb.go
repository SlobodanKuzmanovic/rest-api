package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	Pk_UserId string `json:"pk_UserId"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type HotUser struct {
	Pk_UserId string `json:"pk_UserId"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	CommentCount string `json:"comment_count"`
}
type LogUser struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func UpdateUser (us User) () {
	query := fmt.Sprintf("UPDATE `Users` SET `Name`= '%s', `Surname`= '%s', `Email` = '%s' WHERE Pk_UserId= %s", us.Name, us.Surname, us.Email,us.Pk_UserId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func UpdateUserPass (us User) () {
	hash := md5.Sum([]byte(us.Password))
	hashedPass := hex.EncodeToString(hash[:])
	query := fmt.Sprintf("UPDATE `Users` SET `Password` = '%s' WHERE Pk_UserId= %s", hashedPass,us.Pk_UserId)
	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteUser(id string)()  {
	_, err := DB.Query("DELETE FROM `Users` WHERE Pk_UserId = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func NewUser (us User) () {

	hash := md5.Sum([]byte(us.Password))
	hashedPass := hex.EncodeToString(hash[:])
	query := fmt.Sprintf("INSERT INTO `Users`(`Name`, `Surname`, `Email`, `Password`) VALUES ('%s', '%s', '%s','%s')", us.Name, us.Surname, us.Email, hashedPass)

	_, err := DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func OneUser(email string, password string)(User)  {
	hash := md5.Sum([]byte(password))
	hashedPass := hex.EncodeToString(hash[:])
	results, err := DB.Query("SELECT Name, Surname, Email, Password FROM `users` WHERE Email = ? AND Password =?", email, hashedPass)
	if err != nil {
		panic(err.Error())
	}
	var us User
	for results.Next() {

		err = results.Scan(&us.Name, &us.Surname, &us.Email, &us.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	return us
}


func OneUserById(id string)(User)  {

	results, err := DB.Query("SELECT Pk_UserId, Name, Surname, Email, Password FROM `users` WHERE Pk_UserId = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var us User
	for results.Next() {

		err = results.Scan(&us.Pk_UserId, &us.Name, &us.Surname, &us.Email, &us.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	return us
}


func AllUsers ()([]User, error)  {

	results, err := DB.Query("SELECT Pk_UserId, Name, Surname, Email, Password from Users")
	if err != nil {
		panic(err.Error())
	}

	var Users []User

	for results.Next() {
		var us User

		err = results.Scan(&us.Pk_UserId, &us.Name, &us.Surname, &us.Email, &us.Password)
		if err != nil {
			panic(err.Error())
		}

		Users = append(Users, us)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Users, err
}
func HotUsers ()([]HotUser, error)  {

	results, err := DB.Query("SELECT u.Pk_UserId, u.Name, u.Surname, COUNT(r.Pk_ReplyId) as Hotest " +
		"FROM users AS u JOIN replies AS r ON u.Pk_UserId = r.Fk_UserId " +
		"GROUP BY u.Pk_UserId, u.Name, u.Surname ORDER BY hotest DESC LIMIT 5 OFFSET 0")
	if err != nil {
		panic(err.Error())
	}

	var Users []HotUser

	for results.Next() {
		var us HotUser

		err = results.Scan(&us.Pk_UserId, &us.Name, &us.Surname, &us.CommentCount)
		if err != nil {
			panic(err.Error())
		}

		Users = append(Users, us)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return Users, err
}