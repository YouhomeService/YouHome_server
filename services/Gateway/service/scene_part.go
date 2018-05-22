package service

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

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

func sceneCreateHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	userId := user["userId"].(string)
	sceneName := user["sceneName"].(string)
	fmt.Println(userId + sceneName)
	resp, err := http.Post(sceneService+"/v1/scenes","application/json",bytes.NewBuffer(data))
	checkError(err)
	defer resp.Body.Close()

	body,err:= ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprintf(w, string(body))
	return
}

func deleteSceneHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var user map[string]interface{}
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &user)
	sceneId := user["sceneId"].(string)
	fmt.Println(sceneId)
	resp, err := http.Post(sceneService+"/v1/scenes/delete","application/json",bytes.NewBuffer(data))
	checkError(err)
	defer resp.Body.Close()

	body,err:= ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Fprintf(w, string(body))
	return
}
func sceneNameHandler(w http.ResponseWriter, r*http.Request)  {

}