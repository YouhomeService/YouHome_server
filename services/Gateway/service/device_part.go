package service

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

func deviceHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	method := r.Method
	if method == "GET" {
		getDeviceById(w,r)
	}else if method == "POST"{
		changeDeviceState(w,r)
	}else{
		fmt.Fprint(w,"you can only POST or GET.")
	}
	return
}
func changeDeviceState(w http.ResponseWriter,r * http.Request){
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

	fmt.Fprint(w,string(body))
}
func getDeviceById(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
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
	fmt.Fprint(w, string(body))
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