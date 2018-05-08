package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"bytes"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

var host = "http://127.0.0.1"
var userService = host + ":9091"
var deviceService = host + ":9092"
var sceneService = host + ":9093"

func deviceHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	method := r.Method
	if method == "GET" {
		var deviceId string
		for _,v := range r.Form{
			deviceId = strings.Join(v, "")
		}
		if deviceId == ""{
			fmt.Fprintf(w, "deviceId parse err.")
			return
		}
		fmt.Println("get method will be send.\n")

		resp, err := http.Get(deviceService + "/v1/devices/states?deviceId=" + deviceId)
		checkError(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		checkError(err)

		fmt.Fprintln(w, string(body))
		return
	}else if method == "POST"{
		var deviceId string
		for _,v := range r.Form{
			deviceId = strings.Join(v, "")
		}
		if deviceId == ""{
			fmt.Fprintf(w,"devices parse error.")
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		checkError(err)

		resp, err := http.Post(deviceService+"/v1/devices/states?deviceId="+deviceId,"application/json",bytes.NewBuffer(data))

		checkError(err)


		body, err := ioutil.ReadAll(resp.Body)
		checkError(err)

		fmt.Fprintln(w,string(body))
		return
	}else{
		fmt.Fprintln(w,"you can only POST or GET.")
		return
	}

}
func deviceAllHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var sceneId string
	for k, v := range r.Form{
		if k == "sceneId"{
			sceneId = strings.Join(v, "")
		}
	}
	if sceneId == ""{
		fmt.Fprintf(w,"sceneId parse error.")
		return
	}
	resp ,err := http.Get(deviceService+"/v1/devices?sceneId="+sceneId)
	checkError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprintf(w,string(body))
	return
}
func sceneAllHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var userId string
	for _,v := range r.Form{
		userId = strings.Join(v, "")
	}
	if userId == ""{
		fmt.Fprintf(w,"userId parse error.")
		return
	}
	resp, err := http.Get(sceneService+"/v1/scenes?userId="+userId)
	checkError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprintf(w,string(body))
	return
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET"{
		m, _ := url.ParseQuery(r.URL.RawQuery)
		userId := m["userId"][0]

		resp, err := http.Get(userService+"/v1/users/"+userId)
		checkError(err)

		body, err := ioutil.ReadAll(resp.Body)
		checkError(err)

		fmt.Fprintf(w,string(body))
	}else if r.Method == "POST"{
		var user map[string]interface{}
		data, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(data, &user)
		userId := user["userId"].(string)
		fmt.Println(userId+": "+user["userName"].(string))

		fmt.Println(string(data))

		resp, err := http.Post(userService+"/v1/users?userId="+userId,"application/json",bytes.NewBuffer(data))
		checkError(err)

		body, err := ioutil.ReadAll(resp.Body)
		checkError(err)
		fmt.Fprintf(w, string(body))
	}else{
		fmt.Fprintf(w,"you can only Get or Post.")
	}
	return
}
func checkError(err error) {
	if err != nil{
		panic(err)
	}
}
func main() {
	http.HandleFunc("/scene/all",sceneAllHandler)
	http.HandleFunc("/device/all",deviceAllHandler)
	http.HandleFunc("/device",deviceHandler)
	http.HandleFunc("/user",userHandler)

	//err := http.ListenAndServeTLS(":8080", "cert/cert.pem", "cert/key.pem", nil)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}