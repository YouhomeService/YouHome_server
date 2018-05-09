package main

import (
	"net/http"
	"YouHome-back-end/services/userInfo/service"
	"log"
)

func main() {
	service.LoadRouters()

	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}