package main

import (
	"net/http/httputil"
	"net/http"
	"strings"
	"fmt"
)

var userService = "172.22.16.5:9091"
var deviceService ="172.22.16.4:9092"
var roomService = "172.22.16.3:9093"


/*
var userService = "localhost:9091"
var deviceService ="localhost:9092"
var roomService = "localhost:9093"
 */

func NewMultipleHostsReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		//req.URL.Path = target.Path
		path := req.URL.Path
		service := strings.Split(path,"/")[2]
		//service := data[2]//strings.Split(data[2],"?")[0]
		switch service{
		case "users":
			req.URL.Host = userService
		case "rooms":
			req.URL.Host = roomService
		case "devices":
			req.URL.Host = deviceService
		}

		fmt.Println(req.URL.Scheme)
		fmt.Println(req.URL.Host)
		fmt.Println(req.URL.Path)
		fmt.Println(service)
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := NewMultipleHostsReverseProxy()
	err :=http.ListenAndServeTLS(":443", "license/cert.pem","license/key.pem",proxy)
	if err != nil{
		fmt.Println(err)
	}
}