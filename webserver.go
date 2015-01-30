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
	
	x = strings.Index(ls,"=")
	y = strings.Index(ls,"&")
	email := ls[x+1:y]
	ts := strings.Index(email,"%40")
	email = email[:ts]+"@"+email[ts+3:]
	
	m = ls[y+1:]
	x = strings.Index(m,"=")
	password:= m[x+1:]
	
	_, errcre := createuser.CreateNewuser(username,email,password)
	if errcre != nil {
		fmt.Println(errcre)
		fmt.Fprintf(w, errcre.Error())
	} else{
		fmt.Println(username)
		fmt.Println(email)
		fmt.Println(password)
		fmt.Fprintf(w, "User created")
	}
	
}


func main() {
	http.HandleFunc("/registeruser", index);
	http.ListenAndServe(":8080", nil)
}