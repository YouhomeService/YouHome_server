package main

import (
	"net/http"
	"fmt"
	"github.com/Suenaa/agenda-golang/service/tools"
	"io/ioutil"
)

func main() {


	//entities.Insert("temperature", "sensor.temperature_158d000222c6da", "1")
	//entities.Insert("humidity", "sensor.humidity_158d000222c6da", "1")
	//entities.Insert("illumination", "sensor.illumination_7811dce1bbf3", "1")

	res1, err1 := http.Get("http://127.0.0.1:8080/v1/devices?sceneId=1")
	if err1 == nil {
		fmt.Println("Success")
		defer res1.Body.Close()
	} else {
		tools.Report(err1)
	}

	body1, _ := ioutil.ReadAll(res1.Body)
	fmt.Println(string(body1))

}