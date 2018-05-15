package main

import (
	"log"
	"net/http"
	"YouHome_server/services/Gateway/service"
)

func main() {

	service.InitRoutes()
	err := http.ListenAndServeTLS(":8088", "cert.pem", "key.pem", nil)
	//err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
