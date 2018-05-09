package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	data := struct {
		Id string
		Name string
	}{"1234","tom"}
	buf,err := json.Marshal(data)
	check(err)

	fmt.Println(string(buf))
	var user map[string]interface{}
	json.Unmarshal(buf,&user)
	temp := user["ef"]
	if temp == nil{
		fmt.Println("nil")
	}
}
func check(err error){
	if err != nil{
		panic(err)
	}
}
