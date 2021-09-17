package main

import (
	"database/sql"
	"fmt"
	"rest/models"
	"rest/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	models.DB, err = sql.Open("mysql", "b9025434f58ccf:b9025434f58ccf@/eu-cdbr-west-01.cleardb.com")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully")
	defer models.DB.Close()

	routes.HandleRequests()

}
