package service

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"YouHome_server/services/Scene/entities"
	"net/url"
)

func addScenes(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	userId := user["userId"].(string)
	sceneName := user["sceneName"].(string)
	fmt.Println(sceneName,userId)

	err := entities.AddScene(sceneName, userId)
	checkErr(err)
	sceneId := entities.GetSceneId(sceneName,userId)
	fmt.Fprint(w, sceneId)
	return
}


func deleteScenes(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	sceneId := m["sceneId"][0]
	err := entities.DeleteScene(sceneId)
	checkErr(err)
	fmt.Fprint(w,http.StatusOK)
}


func deleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	fmt.Println(string(data))

	sceneId := user["sceneId"].(string)
	fmt.Println(sceneId)

	err := entities.DeleteScene(sceneId)
	checkErr(err)
	fmt.Fprint(w,http.StatusOK)
	return
}