package service

import (
	"net/http"
	"YouHome_server/services/Device_Service/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadRouters() {

	http.HandleFunc("/v1/devices", deviceInfoHandler)
	http.HandleFunc("/v1/devices/states", statesHandler)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}