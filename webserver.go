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
	http.HandleFunc("/loginuser", loginuser)
	http.HandleFunc("/registercommunity", registercommunity)
	http.HandleFunc("/searchcommunity", searchcommunity)
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

func loginuser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	username := m[x+1:y]
	m = m[y+1:]
	x = strings.Index(m,"=")
	password:= m[x+1:]
	fmt.Println(password,username)
	rows, err := db.Query("SELECT username FROM userlist WHERE( username=? AND password=?)", username, password )
	if err != nil {
		panic(err)
		fmt.Fprintf(w, err.Error())
	} else{
		for rows.Next() {
            var username string
            if err := rows.Scan(&username); err != nil {
							fmt.Println(err)
            }
            fmt.Printf("%s\n", username)
						fmt.Fprintf(w, "User is good")
    }
    if err := rows.Err(); err != nil {
			fmt.Println(err)
    }
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
func searchcommunity(res http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	country := m[x+1:y]
	country = strings.Replace(country, "+", "", -1)
	country = strings.Replace(country, "%28", "(", -1)
	country = strings.Replace(country, "%29", ")", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	state := m[x+1:y]
	state = strings.Replace(state, "+", "", -1)
	state = strings.Replace(state, "%28", "(", -1)
	state = strings.Replace(state, "%29", ")", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	y = strings.Index(m,"&")
	name := m[x+1:y]
	name = strings.Replace(name, "+", "", -1)
	name = strings.Replace(name, "%28", "(", -1)
	name = strings.Replace(name, "%29", ")", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	city := m[x+1:]
	city = strings.Replace(city, "+", "", -1)
	myquery := "SELECT name, id, country, state, city FROM communities WHERE name LIKE '%"+name+"%'  OR country LIKE '%"+country+"%' OR state LIKE '%"+state+"%' OR city LIKE '%"+city+"%'";
	rows, err := db.Query(myquery)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
	}
	defer rows.Close()
	id := 0
	for rows.Next() {
		err := rows.Scan(&name, &id, &country, &state, &city)
		if err != nil {
			fmt.Println(err)
      fmt.Fprintln(res,err)
		}
		fmt.Fprintln(res, id, name, country, state, city, "@")
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
    fmt.Fprintln(res, err)
	}
}
