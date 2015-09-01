package main

import (
	"net/http"
	"fmt"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"github.com/melvinodsa/JyApp/users"
	"github.com/melvinodsa/JyApp/communities"
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
	http.HandleFunc("/loginuser", loginuser)
	http.HandleFunc("/registercommunity", registercommunity)
	http.HandleFunc("/searchcommunity", searchcommunity)
	http.ListenAndServe(":8080", nil)
}

func registeruser(w http.ResponseWriter, r *http.Request) {
	users.RegisterUser(w, r, db)
}

func loginuser(w http.ResponseWriter, r *http.Request) {
	users.LoginUser(w, r, db)
}


func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world Melvinodsa")
}

func registercommunity(res http.ResponseWriter, req *http.Request) {
	communities.RegisterCommunity(res, req, db)
}

func searchcommunity(res http.ResponseWriter, req *http.Request) {
	communities.SearchCommunity(res, req, db)
}
