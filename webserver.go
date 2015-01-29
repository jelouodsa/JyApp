package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/jelouodsa/JyApp/createuser"
)


func index(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	m := string(body)
	x := strings.Index(m,"=")
	y := strings.Index(m,"&")
	username := m[x+1:y]
	ls := m[y+1:]
	fmt.Println(username)
	x = strings.Index(ls,"=")
	y = strings.Index(ls,"&")
	email := ls[x+1:y]
	ts := strings.Index(email,"%40")
	email = email[:ts]+"@"+email[ts+3:]
	fmt.Println(email)
	m = ls[y+1:]
	x = strings.Index(m,"=")
	password:= m[x+1:]
	fmt.Println(password)
	_, _ = createuser.CreateNewuser(username,email,password)
	fmt.Fprintf(w, "0")
}


func main() {
	http.HandleFunc("/", index);
	http.ListenAndServe(":8080", nil)
}