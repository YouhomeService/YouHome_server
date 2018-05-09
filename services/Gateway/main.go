package main

import (
	"log"
	"net/http"
	"YouHome_server/services/Gateway/service"
)

func main() {

	service.InitRoutes()
	//err := http.ListenAndServeTLS(":8080", "license/license.pem", "license/key.pem", nil)
	err := http.ListenAndServe(":8088",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}