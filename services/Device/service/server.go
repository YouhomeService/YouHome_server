package service

import (
	"net/http"
	"YouHome_server/services/Device/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadRouters() {

	http.HandleFunc("/v1/devices", deviceInfoHandler)
	http.HandleFunc("/v1/devices/states", statesHandler)
	http.HandleFunc("/v1/devices/devicename", devicenameHandler)
}

func deviceInfoHandler(w http.ResponseWriter,req *http.Request) {

	fmt.Println("deviceInfoHandler")
	if req.Method == "GET" {
		getDeviceHandler(w, req)
	}
}

func statesHandler(w http.ResponseWriter,req *http.Request) {
	if req.Method == "GET" {
		getStatesHandler(w, req)
	}
}

func devicenameHandler(w http.ResponseWriter,req *http.Request) {
	if req.Method == "GET" {
		getDeviceNameHandler(w, req)
	}
	if req.Method == "POST" {
		postDeviceNameHandler(w, req)
	}
}

func getDeviceHandler(w http.ResponseWriter,req *http.Request) {
	id := req.FormValue("sceneId")
	devices := entities.GetDevicesBySceneId(id)
	l := len(devices)
	type device struct {
		DeviceId string `json:"deviceId"`
		DeviceName string `json:"deviceName"`
		EntityId string `json:"entityId"`
	}
	result := make([]device, 0)
	for i := 0; i < l; i++ {
		temp := device{devices[i][0], devices[i][1], devices[i][2]}
		result = append(result, temp)
	}
	data, _ := json.Marshal(result)
	fmt.Fprintln(w, string(data))
}

func getStatesHandler(w http.ResponseWriter,req *http.Request) {
	id := req.FormValue("deviceId")
	entityId := entities.GetEntityId(id)
	haIP := "haserverip"
	res, err := http.Get(haIP+"/api/states/"+entityId)
	check(err)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintln(w, string(body))
}

func getDeviceNameHandler(w http.ResponseWriter,req *http.Request) {
	id := req.FormValue("deviceId")
	deviceName := entities.GetDeviceName(id)
	data := struct{
		DeviceName string `json:"deviceName"`
	}{deviceName}
	buf, _ := json.Marshal(data)
	//fmt.Println(string(buf))
	fmt.Fprintln(w, string(buf))
}

func postDeviceNameHandler(w http.ResponseWriter,req *http.Request) {
	req.ParseForm()
	var device map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &device)
	id := device["deviceId"].(string)
	name := device["deviceName"].(string)
	err := entities.SetDeviceName(id, name)
	check(err)

	if err == nil {
		name := entities.GetDeviceName(id)
		data := struct{
			DeviceName string `json:"deviceName"`
		}{name}
		buf, _ := json.Marshal(data)
		fmt.Fprintln(w, string(buf))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}