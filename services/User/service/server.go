package service

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"YouHome_server/services/User/entities"
)

func LoadRouters() {

	http.HandleFunc("/v1/users", userInfoHandler)
	http.HandleFunc("/v1/users/userName", userNameHandler)
	http.HandleFunc("/v1/test", testHandler)
}

func testHandler(w http.ResponseWriter,req *http.Request) {
	test := "testing!"
	fmt.Fprintln(w, test)
}

func userInfoHandler(w http.ResponseWriter,req *http.Request) {

	if req.Method == "POST" {
		postUserInfoHandler(w, req)
	}

	if req.Method == "GET" {
		getUserName(w, req)
	}
}

func userNameHandler(w http.ResponseWriter,req *http.Request) {

	req.ParseForm()
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &user)
	id := user["userId"].(string)
	name := user["userName"].(string)
	err := entities.SetUserName(id, name)
	check(err)

	if err == nil {
		name := entities.GetNameById(id)
		data := struct{
			UserName string `json:"userName"`
		}{name}
		buf, _ := json.Marshal(data)
		fmt.Fprintln(w, string(buf))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func postUserInfoHandler(w http.ResponseWriter,req *http.Request) {

	req.ParseForm()
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &user)
	code := user["code"].(string)
	appId := "wx394f9cc0f949d50b"
	secret := "0346967483b301c189254bef576b1091"
	getUrl := "https://api.weixin.qq.com/sns/jscode2session?appid="+appId+"&secret="+secret+"&js_code="+code+"&grant_type=authorization_code"
	res, err := http.Get(getUrl)
	var u map[string]interface{}
	body1, err1 := ioutil.ReadAll(res.Body)
	check(err1)
	json.Unmarshal(body1, &u)
	errCode := u["errcode"]

	if errCode != nil {
		data := struct{
			UserId string `json:"userId"`
		}{"0"}
		buf, _ := json.Marshal(data)
		fmt.Fprintln(w, string(buf))
		return
	}

	openId := u["openid"].(string)

	err = entities.Login(openId)

	if err == nil {
		data := struct{
			UserId string `json:"userId"`
		}{openId}
		buf, _ := json.Marshal(data)
		fmt.Fprintln(w, string(buf))

	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func getUserName(w http.ResponseWriter,req *http.Request) {

	id := req.FormValue("userId")
	name := entities.GetUserName(id)
	data := struct{
		UserName string `json:"userName"`
	}{name}
	buf, _ := json.Marshal(data)
	fmt.Println(string(buf))
	fmt.Fprintln(w, string(buf))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
