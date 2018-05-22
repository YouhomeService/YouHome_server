package service

import (
	"net/http"
	"net/url"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"YouHome_server/services/Scene/entities"
)


func handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if r.Method == "GET"{
		getScenes(w,r)
	} else if r.Method == "POST"{
		addScenes(w,r)
	} else if r.Method =="DELETE"{
		deleteScenes(w,r)
	}
	return
}
func getScenes(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	m, _ := url.ParseQuery(r.URL.RawQuery)
	userId := m["userId"][0]
	data := entities.GetScenes(userId)
	fmt.Fprint(w, data)
	return
}
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
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
func LoadRoute() {
	http.HandleFunc("/v1/scenes/delete",deleteHandler)
	http.HandleFunc("/v1/scenes", handler)
}
