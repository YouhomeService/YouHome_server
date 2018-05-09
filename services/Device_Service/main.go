package main

import (
	"net/http"
	"YouHome-back-end/services/deviceInfo/service"
	"log"
)

func main() {
	service.LoadRouters()

	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}