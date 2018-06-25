package main

import (
	"net/http/httputil"
	"net/http"
	"strings"
	"fmt"
)
var voiceService = "172.22.16.6:3005"

var userService = "172.22.16.5:9091"
var deviceService ="172.22.16.4:9092"
var roomService = "172.22.16.3:9093"


/*
var userService = "localhost:9091"
var deviceService ="localhost:9092"
var roomService = "localhost:9093"
 */

func checkLogin (req *http.Request) bool {
	headers := req.Header;
	value, existed := headers["Authorization"]
	if (existed && len(value) != 0 && value[0] != "") {
		return true;
	}
	return false;
}

func NewMultipleHostsReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		//hasLogged := checkLogin()
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
		case "voice":
			req.URL.Host = voiceService
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
	err :=http.ListenAndServeTLS(":443", "license/_youhome.xyz_bundle.crt","license/2_youhome.xyz.key",proxy)
	if err != nil{
		fmt.Println(err)
	}
}