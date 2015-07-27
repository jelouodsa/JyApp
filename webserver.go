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
	err error
)

func main() {
	db, err = sql.Open("mysql", "root:q@tcp(localhost:3306)/candlelight")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/", hello)
	http.HandleFunc("/registeruser", registeruser)
	http.HandleFunc("/registercommunity", registercommunity)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world Melvinodsa")
}

func registeruser(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(email,password,username)
	_, err := db.Query("INSERT INTO userlist VALUES( ?, ?, ? )", username, email, password )
	if err != nil {
		panic(err)
		fmt.Fprintf(w, err.Error())
	} else{
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Fprintf(w, "User created")
	}
}

func registercommunity(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	privacy := m[x+1:y]
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	country := m[x+1:y]
	country = strings.Replace(country, "+", " ", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	state := m[x+1:y]
	state = strings.Replace(state, "+", " ", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	name := m[x+1:y]
	name = strings.Replace(name, "+", " ", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	city := m[x+1:y]
	city = strings.Replace(city, "+", " ", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	admin := m[x+1:]
	admin = strings.Replace(admin, "+", " ", -1)
	_, err := db.Query("INSERT INTO communities (name,admin,privacy,country,state,city) VALUES( ?, ?, ?, ?, ?, ? )", name, admin, privacy, country, state, city )
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
	} else{
		fmt.Println(privacy,country,state,name,city,admin)
		fmt.Fprintf(res, "Community created")
	}
}
