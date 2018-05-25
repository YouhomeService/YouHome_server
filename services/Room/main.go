package main

import (
	"net/http"
	"log"
	"fmt"
	"YouHome_server/services/Room/service"
)

func main() {
	service.LoadRoute()

	err := http.ListenAndServe(":9093", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}else{
		fmt.Println("working...")
	}
}