package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func main() {


	//entities.Insert("temperature", "sensor.temperature_158d000222c6da", "1")
	//entities.Insert("humidity", "sensor.humidity_158d000222c6da", "1")
	//entities.Insert("illumination", "sensor.illumination_7811dce1bbf3", "1")

	//res1, err1 := http.Get("http://127.0.0.1:9092/v1/devices/devicename?deviceId=1")

	data := struct {
		DeviceId string `json:"deviceId"`
		DeviceName string `json:"deviceName"`
	}{"1","anothertemperature"}
	buf, _ := json.Marshal(data)
	res1, err1 := http.Post("http://127.0.0.1:9092/v1/devices/devicename",
		"application/json", bytes.NewBuffer(buf))

	if err1 == nil {
		fmt.Println("Success")
		defer res1.Body.Close()
	} else {
		//tools.Report(err1)
	}

	body1, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body1))

}