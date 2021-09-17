package main

import (
	"database/sql"
	"fmt"
	"log"
	"rest/models"
	"rest/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	log.Printf("Try open Connection")
	models.DB, err = sql.Open("mysql", "b9025434f58ccf:28bc132c@tcp(eu-cdbr-west-01.cleardb.com:3306)/heroku_ecd70f5db7afe86")
	log.Printf("BBBBBBBBBBBBBBBBBB")
	if err != nil {
		log.Printf(err.Error())

		panic(err.Error())
	}
	log.Printf("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	fmt.Println(models.DB.Stats())
	log.Printf("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	defer models.DB.Close()

	routes.HandleRequests()

}
