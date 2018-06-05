package service

import (
	"net/http"
	"YouHome_server/services/Device/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
	"strings"
	"net/url"
	"github.com/tidwall/gjson"
)

func LoadRouters() {

	http.HandleFunc("/v1/devices/available",availableHandler)
	http.HandleFunc("/v1/devices/states", statesHandler)
	http.HandleFunc("/v1/devices/devicename", devicenameHandler)
	http.HandleFunc("/v1/devices/url",urlHandler)
	http.HandleFunc("/v1/devices", deviceInfoHandler)
}

var HAaddr = "http://123.207.55.27:8125"

func deviceInfoHandler(w http.ResponseWriter,req *http.Request) {

	fmt.Println("deviceInfoHandler")
	if req.Method == "GET" {
		getDeviceHandler(w, req)
	}
	if req.Method == "POST" {
		postDeviceHandler(w, req)
	}
}

func urlHandler(w http.ResponseWriter,req *http.Request){
	if req.Method == "POST"{
		updateUrl(w,req)
	}else{
		getUrl(w,req)
	}
}

func statesHandler(w http.ResponseWriter,req *http.Request) {
	if req.Method == "GET" {
		getStatesHandler(w, req)
	}
	if req.Method == "POST" {
		postStatesHandler(w, req)
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
	id := req.FormValue("roomId")
	devices := entities.GetDevicesByRoomId(id)
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

func postDeviceHandler(w http.ResponseWriter,req *http.Request) {
	req.ParseForm()
	var device map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &device)
	eid := device["entityId"].(string)
	name := device["deviceName"].(string)
	rid := device["roomId"].(string)

	err, deviceid := entities.AddDevice(name, eid, rid)
	check(err)
	if err == nil {
		data := struct{
			DeviceId string `json:"deviceId"`
		}{deviceid}
		buf, _ := json.Marshal(data)
		fmt.Fprintln(w, string(buf))
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func getStatesHandler(w http.ResponseWriter,req *http.Request) {
	id := req.FormValue("deviceId")
	entityId := entities.GetEntityId(id)
	haIP := "http://123.207.55.27:8125"
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

func postStatesHandler(w http.ResponseWriter,req *http.Request) {
	req.ParseForm()
	var device map[string]interface{}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &device)
	id := device["deviceId"].(string)
	operation := device["operation"].(string)
	entityId := entities.GetEntityId(id)
	data := struct {
		EntityId string `json:"entity_id"`
	}{entityId}
	domain := strings.Split(entityId, ".")
	buf, _ := json.Marshal(data)
	fmt.Println(HAaddr+"/api/services/"+domain[0]+"/"+operation)
	_, err := http.Post(HAaddr+"/api/services/"+domain[0]+"/"+operation,
		"application/json", bytes.NewBuffer(buf))
	var r string
	if err == nil {
		r = "1"
	} else {
		r = "0"
	}
	rdata := struct {
		Result string `json:"result"`
	}{r}
	buf1, _ := json.Marshal(rdata)
	fmt.Fprintln(w, string(buf1))
}

func updateUrl(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	deviceId := user["deviceId"].(string)
	deviceUrl := user["url"].(string)
	result := entities.UpdateDeviceUrl(deviceId,deviceUrl)
	fmt.Fprint(w,result)
}

func getUrl(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	deviceId := m["deviceId"][0]
	data := entities.GetDeviceUrl(deviceId)
	fmt.Fprint(w, data)
	return
}

func availableHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	resp, _ := http.Get(HAaddr+"/api/states")
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	jsonStream := string(body)

	var re []string
	result := gjson.Get(jsonStream, "#.entity_id")
	for _, name := range result.Array() {
		//println(name.String())
		id := name.String()
		temp :=strings.Split(id,".")
		if temp[0] == "switch" || temp[0] == "light" || temp[0]=="sensor"{
			//println(name.String())
			re = append(re, name.String())
		}
	}
	data, _ :=json.Marshal(re)
	fmt.Fprint(w, string(data))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}