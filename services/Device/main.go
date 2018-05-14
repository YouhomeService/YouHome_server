package main

import (
	"net/http"
	"log"
	"YouHome_server/services/Device/service"
)

func main() {
	service.LoadRouters()

	err := http.ListenAndServe(":9092",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}