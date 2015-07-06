package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world Melvinodsa")
}

func index(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	username := m[x+1:y]
	ls := m[y+1:]

	x = strings.Index(ls,"=")
	y = strings.Index(ls,"&")
	email := ls[x+1:y]
	ts := strings.Index(email,"%40")
	email = email[:ts]+"@"+email[ts+3:]

	m = ls[y+1:]
	x = strings.Index(m,"=")
	password:= m[x+1:]

	_, err := db.Query("INSERT INTO userlist VALUES( ?, ?, ? )", username, email, password )
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
	} else{
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Fprintf(w, "User created")
	}
}



func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/registeruser", index);
	http.ListenAndServe(":8080", nil)
	db, err = sql.Open("mysql", "root:q@tcp(localhost:3306)/candlelight")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
