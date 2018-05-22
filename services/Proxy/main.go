package main

import (
	"net/http/httputil"
	"net/http"
	"strings"
	"fmt"
)

/*
var userService = "http://172.22.16.5:9091"
var deviceService ="http://172.22.16.4:9092"
var sceneService = "http://172.22.16.3:9093"
 */
var userService = "localhost:9091"
var deviceService ="localhost:9092"
var sceneService = "localhost:9093"
func NewMultipleHostsReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		//req.URL.Path = target.Path
		path := req.URL.Path
		data := strings.Split(path,"/")
		service := strings.Split(data[2],"?")[0]
		switch service{
		case "users":
			req.URL.Host = userService
		case "scenes":
			req.URL.Host = sceneService
		case "devices":
			req.URL.Host = deviceService
		}

		fmt.Println(req.URL.Scheme)
		fmt.Println(req.URL.Host)
		fmt.Println(req.URL.Path)
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	//err := http.ListenAndServeTLS(":8088", "cert.pem", "key.pem", nil)

	proxy := NewMultipleHostsReverseProxy()
	err :=http.ListenAndServeTLS(":443", "cert.pem","key.pem",proxy)
	if err != nil{
		fmt.Println(err)
	}
}