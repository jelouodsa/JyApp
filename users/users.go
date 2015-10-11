package users

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "strings"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var err error

func RegisterUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body, _ := ioutil.ReadAll(r.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	username := m[x+1:y]
	username = strings.Replace(username, "+", " ", -1)
	username = strings.Replace(username, "%28", "(", -1)
	username = strings.Replace(username, "%29", ")", -1)
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

func LoginUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body, _ := ioutil.ReadAll(r.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	username := m[x+1:y]
	username = strings.Replace(username, "+", " ", -1)
	username = strings.Replace(username, "%28", "(", -1)
	username = strings.Replace(username, "%29", ")", -1)
	m = m[y+1:]
	x = strings.Index(m,"=")
	password:= m[x+1:]
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
