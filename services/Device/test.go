package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {


	//entities.Insert("temperature", "sensor.temperature_158d000222c6da", "1")
	//entities.Insert("humidity", "sensor.humidity_158d000222c6da", "1")
	//entities.Insert("light", "light.gateway_light_7811dce1bbf3", "1")

	//entityId := "light.gateway_light_7811dce1bbf3"
	//res1, err1 := http.Get("http://127.0.0.1:9092/v1/devices?sceneId=1")
	//res1, err1 := http.Get("http://127.0.0.1:9092/v1/devices/states?deviceId=4")

	data := struct {
		DeviceId string `json:"deviceId"`
	} {"15"}
	buf, _ := json.Marshal(data)
	res1, err1 := http.Post("http://127.0.0.1:9092/v1/devices/delete",
		"application/json", bytes.NewBuffer(buf))
	if err1 == nil {
		fmt.Println("Success")
		defer res1.Body.Close()
	} else {
		//tools.Report(err1)
		fmt.Println(err1)
	}
	body, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body))
/*
	data := struct {
		EntityId string `json:"entityId"`
		DeviceName string `json:"deviceName"`
		RoomId string `json:"roomId"`
		Url string `json:"url"`
	}{"testentityid","testname", "1", "testurl"}
	buf, _ := json.Marshal(data)
	res1, err1 := http.Post("http://127.0.0.1:9092/v1/devices",
		"application/json", bytes.NewBuffer(buf))

	if err1 == nil {
		fmt.Println("Success")
		defer res1.Body.Close()
	} else {
		//tools.Report(err1)
		fmt.Println(err1)
	}
	body, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body))
	 // /v1/devices?sceneId=abc
*/
	res1, _ = http.Get("http://127.0.0.1:9092/v1/devices?roomId=1")

	body1, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body1))

}