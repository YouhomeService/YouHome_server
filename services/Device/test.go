package main

import (
	"fmt"
	"io/ioutil"
	//"encoding/json"
	"net/http"
	//"bytes"
)

func main() {


	//entities.Insert("temperature", "sensor.temperature_158d000222c6da", "1")
	//entities.Insert("humidity", "sensor.humidity_158d000222c6da", "1")
	//entities.Insert("light", "light.gateway_light_7811dce1bbf3", "1")

	//entityId := "light.gateway_light_7811dce1bbf3"
	//res1, err1 := http.Get("http://127.0.0.1:9092/v1/devices?sceneId=1")
	//res1, err1 := http.Get("http://127.0.0.1:9092/v1/devices/states?deviceId=4")


	/*
	data := struct {
		DeviceId string `json:"deviceId"`
		Operation string `json:"operation"`
	}{"4","turn_off"}
	buf, _ := json.Marshal(data)
	res1, err1 := http.Post("http://127.0.0.1:9092/v1/devices/states",
		"application/json", bytes.NewBuffer(buf))

	if err1 == nil {
		fmt.Println("Success")
		defer res1.Body.Close()
	} else {
		//tools.Report(err1)
	}
	 */
	 // /v1/devices?sceneId=abc
	res1, _ := http.Get("http://127.0.0.1:9092/v1/devices/states?deviceId=4")

	body1, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body1))

}