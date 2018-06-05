package main

import (
	"net/http"
	"YouHome_server/services/Voice/service"
	"log"
)
 
func main() {
	service.LoadRouters()

	err := http.ListenAndServe(":9094",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}