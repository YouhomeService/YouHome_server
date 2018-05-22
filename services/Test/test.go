package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
	"bytes"
	"encoding/json"
	"crypto/tls"
)

var address string = "https://123.207.55.27:8088"
var tr = &http.Transport{
TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
}
var client = &http.Client{Transport: tr}
func deleteScene(id string){
	data := struct {
		SceneId string `json:"sceneId"`
	}{id}
	buf , _ := json.Marshal(data)///scene/delete
	//fmt.Println(string(buf))
	//resp, err = http.Post(address + "/scene/delete","application/json",bytes.NewBuffer(tempJson))
	resp, err := client.Post(address+"/scene/delete","application/json",bytes.NewBuffer(buf))
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func login(code string){
	user := struct{
		Code string `json:"code"`
	}{code}
	//fmt.Println(string(user))
	userJson ,err := json.Marshal(user)
	fmt.Println(string(userJson))
	checkErr(err)
	resp, err := client.Post(address+"/user/login","application/json",bytes.NewBuffer(userJson))
	checkErr(err)
	body,err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func createScene(userId,sceneName string){
	scene := struct {
		UserId string `json:"userId"`
		SceneName string `json:"sceneName"`
	}{userId,sceneName}
	sceneJson,err := json.Marshal(scene)
	checkErr(err)
	resp, err := client.Post(address+"/scene/create","application/json",bytes.NewBuffer(sceneJson))
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getAllScene(userId string){
	resp, err := client.Get(address + "/scene/all?userId="+userId)
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getUserInfo(userId string){
	resp, err := client.Get(address+"/user?userId="+userId)
	checkErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func getDeviceOfScene(sceneId string){
	//device/all?sceneId=room
	res,err := client.Get(address + "/device/all?sceneId=" + sceneId)
	checkErr(err)

	body,err := ioutil.ReadAll(res.Body)
	checkErr(err)

	fmt.Println(string(body))
}
func updateUserInfo(userId,userName string){
	temp := struct {
		UserId string `json:"userId"`
		UserName string `json:"userName"`
	}{userId,userName}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := client.Post(address + "/user","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
func TestWechatApi(){
	//https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wx394f9cc0f949d50b&secret=0346967483b301c189254bef576b1091&js_code=JSCODE&grant_type=authorization_code")
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
}
func main() {


	res,err:= client.Get( "http://localhost:8088/v1/users?userId=1533")
	checkErr(err)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))



	/*
	temp := struct {
		SceneId string `json:"sceneId"`
		SceneName string `json:"sceneName"`
	}{"1","bedroom"}
	buf,err := json.Marshal(temp)
	checkErr(err)
	resp, err := client.Post("http://localhost:8088/v1/scenes/name","application/json",bytes.NewBuffer(buf))
	checkErr(err)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	 */


}




