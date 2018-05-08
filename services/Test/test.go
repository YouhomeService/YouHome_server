package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
)


func main() {
	data := struct{
		UserId string `json:"userId"`
		UserName string `json:"userName"`
	}{"abc","Jack"}
	buf, _ := json.Marshal(data)
	// /v1/users/userName?userId=abc
	resp, err := http.Post("http://localhost:8080/user","application/json",bytes.NewBuffer(buf))

    if err != nil{
    	fmt.Println(err)
    	return
    }
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
    	fmt.Println(err)
    }
    fmt.Println("--"+string(body)+"++")
}



