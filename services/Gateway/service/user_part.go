package service

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"bytes"
	"net/url"
	"encoding/json"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET"{
		getUserInfo(w,r)
	}else if r.Method == "POST"{
		changeUserInfo(w,r)
	}else{
		fmt.Fprintf(w,"you can only Get or Post.")
	}
	return
}
func getUserInfo(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	userId := m["userId"][0]

	resp, err := http.Get(userService+"/v1/users?userId="+userId)
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprintf(w,string(body))
}
func changeUserInfo(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	userId := user["userId"].(string)
	fmt.Println(userId+": "+user["userName"].(string))

	fmt.Println(string(data))


	resp, err := http.Post(userService+"/v1/users/userName","application/json",bytes.NewBuffer(data))
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Fprintf(w, string(body))
}
func userLoginHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	data , err:= ioutil.ReadAll(r.Body)
	checkError(err)

	resp, err := http.Post(userService+"/v1/users","application/json",bytes.NewBuffer(data))
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprint(w, string(body))

}