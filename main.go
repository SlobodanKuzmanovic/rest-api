package main

import (
	"database/sql"
	"fmt"
	"rest/models"
	"rest/routes"
	"io"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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

	http.HandleFunction("/", hello)
}



func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Hello World!")
}