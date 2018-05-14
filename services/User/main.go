package main

import (
	"net/http"
	"YouHome_server/services/User/service"
	"log"
)

func main() {
	service.LoadRouters()

	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}