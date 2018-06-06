package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"YouHome_server/services/Room/entities"
	"net/url"
)

func addRooms(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	userId := user["userId"].(string)
	roomName := user["roomName"].(string)
	roomUrl := user["url"].(string)
	fmt.Println(roomName,userId,roomUrl)

	err := entities.AddRoom(roomName, userId,roomUrl)
	checkErr(err)
	roomId := entities.GetRoomId(roomName,userId)
	fmt.Fprint(w, roomId)
	return
}


func deleteRoom(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	roomId := m["roomId"][0]
	err := entities.DeleteRoom(roomId)
	checkErr(err)
	fmt.Fprint(w,http.StatusOK)
}


func deleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	roomId := user["roomId"].(string)
	fmt.Println(roomId)

	err := entities.DeleteRoom(roomId)
	checkErr(err)
	fmt.Fprint(w,http.StatusOK)
	return
}