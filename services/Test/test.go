package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
	"bytes"
	"encoding/json"
)

var address string = "http://localhost:8088"
func deleteScene(id string){
	data := struct {
		SceneId string `json:"sceneId"`
	}{id}
	buf , _ := json.Marshal(data)///scene/delete
	//fmt.Println(string(buf))
	//resp, err = http.Post(address + "/scene/delete","application/json",bytes.NewBuffer(tempJson))
	resp, err := http.Post(address+"/scene/delete","application/json",bytes.NewBuffer(buf))
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func login(){
	user := struct{
		UserId string `json:"userId"`
	}{"0613"}
	userJson ,err := json.Marshal(user)
	checkErr(err)
	resp, err := http.Post(address+"/user/login","application/json",bytes.NewBuffer(userJson))
	checkErr(err)
	body,err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func createScene(){
	scene := struct {
		UserId string `json:"userId"`
		SceneName string `json:"sceneName"`
	}{"0613","room"}
	sceneJson,err := json.Marshal(scene)
	checkErr(err)
	resp, err := http.Post(address+"/scene/create","application/json",bytes.NewBuffer(sceneJson))
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getAllScene(){
	resp, err := http.Get(address + "/scene/all?userId=0613")
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getUserInfo(userId string){
	resp, err := http.Get(address+"/user?userId="+userId)
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func changeUserInfo(){
	temp := struct {
		UserId string `json:"userId"`
		UserName string `json:"userName"`
	}{"1533","liangfeng"}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := http.Post(address + "/user","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {

/*
temp := struct{
		SceneId string `json:"sceneId"`
	}{"7"}
	tempJson ,_:= json.Marshal(temp)
	resp, err = http.Post(address + "/scene/delete","application/json",bytes.NewBuffer(tempJson))
	checkErr(err)
	body,_ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
 */
 	changeUserInfo()
 	getUserInfo("1533")
	//getAllScene()
	//√deleteScene("7")
}
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}



