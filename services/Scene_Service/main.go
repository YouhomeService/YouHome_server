package main

import (
	"Youhome/services/Scene/service"
	"net/http"
	"log"
	"fmt"
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