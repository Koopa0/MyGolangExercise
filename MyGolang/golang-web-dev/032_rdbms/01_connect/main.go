package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main()  {
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(mydbinstance.cakwl95bxza0.us-west-1.rds.amazomaws.com:3306)/koopa?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}
func index(w http.ResponseWriter, req *http.Request)  {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}
func check(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}
