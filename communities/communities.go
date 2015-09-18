package communities

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var err error

func RegisterCommunity(res http.ResponseWriter, req *http.Request, db *sql.DB) {
  id := 0
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
	doneFlag := false
	rows, err := db.Query("SELECT username FROM userlist WHERE( username=?)", admin )
	if err != nil {
		panic(err)
		fmt.Fprintf(res, err.Error())
	} else{
		if rows.Next() {
			_, err = db.Query("INSERT INTO communities (name,admin,privacy,country,state,city) VALUES( ?, ?, ?, ?, ?, ? )", name, admin, privacy, country, state, city )
      myquery := "SELECT id FROM communities WHERE username=?";
      rows, err := db.Query(myquery,admin)
    	if err != nil {
    		fmt.Println(err)
    		fmt.Fprintf(res, err.Error())
    	}
    	defer rows.Close()
    	for rows.Next() {
    		err := rows.Scan(&id)
    		if err != nil {
    			fmt.Println(err)
          fmt.Fprintln(res,err)
    		}
    	}
    	err = rows.Err()
    	if err != nil {
    		fmt.Println(err)
        fmt.Fprintln(res, err)
    	}
      _, err = db.Query("INSERT INTO communitymember (id,privilage,username) VALUES( ?, ?, ? )", id, 0, admin )
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(res, err.Error()+"$")
			} else{
				doneFlag = true
			}
    }
    if err := rows.Err(); err != nil {
			fmt.Println(err)
    }
	}
	if doneFlag {
    x := strconv.Itoa(id)
    fmt.Fprintf(res, "Community created$"+x)
	}
}

func JoinCommunity(res http.ResponseWriter, req *http.Request, db *sql.DB) {
  body, _ := ioutil.ReadAll(req.Body)
  m := string(body)
  x := strings.Index(m,"=")
  y := strings.Index(m,"&")
  id := m[x+1:y]
  m = m[y+1:]
  x = strings.Index(m,"=")
	username := m[x+1:]
  _, err := db.Query("INSERT INTO communitymember (id,privilage,username) VALUES( ?, ?, ? )", id, 1, username )
  if err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
	} else {
    fmt.Fprintln(res, "Joined community")
  }
}

func FollowCommunity(res http.ResponseWriter, req *http.Request, db *sql.DB) {
  body, _ := ioutil.ReadAll(req.Body)
  m := string(body)
  x := strings.Index(m,"=")
  y := strings.Index(m,"&")
  id := m[x+1:y]
  m = m[y+1:]
  x = strings.Index(m,"=")
	username := m[x+1:]
  _, err := db.Query("INSERT INTO communitymember (id,privilage,username) VALUES( ?, ?, ? )", id, 2, username )
  if err != nil {
		fmt.Println(err)
		fmt.Fprintf(res, err.Error())
	} else {
    fmt.Fprintln(res, "Followed community")
  }
}

func SearchCommunity(res http.ResponseWriter, req *http.Request, db *sql.DB) {
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
	myquery := "SELECT name, id, country, state, city FROM communities WHERE privacy =0 AND name LIKE '%"+name+"%'  AND country LIKE '%"+country+"%' AND state LIKE '%"+state+"%' AND city LIKE '%"+city+"%'";
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
