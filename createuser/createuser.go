package createuser

import (
	"fmt"
	"os"
	"encoding/xml"
)

type Account struct {
	XMLName xml.Name `xml:"account"`
	Id int `xml:"id"`
	Username string `xml:"username"`
	Email string `xml:"email"`
	Password string `xml:"password"`
}

func updateUserEntry(user Account) (error) {
	file, err := os.Create("users/"+string(user.Username))
	if err != nil {
		return err
	}
	defer file.Close()
	output, err1 := xml.MarshalIndent(user, "  ", "    ")
    if err1 != nil {
    	return  err1
    }
	if _, err = file.WriteString(string(output)); err != nil {
		return  err
	}
	return nil
}

func getUserId(off int) (int, error) {
	f, err := os.Open("userid.us")
	if err != nil {
		return -1, err
	}
	defer f.Close()
	var num int
	_, err = fmt.Fscanf(f,"%d",&num)
	if err != nil {
		return -2, err
	}
	f.Close()
	f, err = os.Create("userid.us")
	if err != nil {
		return -3, err
	}
	num = num + off
	_, err = fmt.Fprintf(f,"%d",num)
	if err != nil {
		return -4, err      
    }
	return num, nil
}


func CreateNewuser(username string, email string, password string) (int, error) {	
	ids, err := getUserId(1)
	if err != nil{
		return -1, err
	}
	v := Account{Id : ids,Username : username,Email : email,Password : password}
	if err = updateUserEntry(v); err != nil{
		_,_ = getUserId(-1)
		return -2, err
	}
	return 0, nil
}