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
	models.DB, err = sql.Open("mysql", "b9025434f58ccf:b9025434f58ccf@/eu-cdbr-west-01.cleardb.com")
	log.Printf("BBBBBBBBBBBBBBBBBB")
	if err != nil {
		log.Printf(err.Error())

		panic(err.Error())
	}
	log.Printf("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	fmt.Println("Successfully")
	defer models.DB.Close()

	routes.HandleRequests()

}
