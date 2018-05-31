package service

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"YouHome_server/services/Room/entities"
)

func getRooms(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	userId := m["userId"][0]
	data := entities.GetRooms(userId)
	fmt.Fprint(w, data)
	return
}
func nameHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if r.Method == "GET"{
		getName(w,r);
	}else{
		updateName(w,r)
	}
	return
}
func getName(w http.ResponseWriter,r*http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	roomId := m["roomId"][0]
	data := entities.GetRoomName(roomId)
	fmt.Fprint(w, data)
	return
}
func updateName(w http.ResponseWriter,r*http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	roomId := user["roomId"].(string)
	roomName := user["roomName"].(string)
	result := entities.UpdateRoomName(roomName,roomId)
	fmt.Fprint(w,result)
}

func getRoomUrl(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	roomId := m["roomId"][0]
	data := entities.GetRoomUrl(roomId)
	fmt.Fprint(w, data)
	return
}

func updateRoomUrl(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	roomId := user["roomId"].(string)
	roomUrl := user["url"].(string)
	result := entities.UpdateRoomUrl(roomUrl,roomId)
	fmt.Fprint(w,result)
}


