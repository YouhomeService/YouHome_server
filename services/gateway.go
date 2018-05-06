package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"net/url"
	"io/ioutil"
)

var host = "127.0.0.1"
var userService = host + ":9091"
var deviceService = host + ":9092"
var sceneService = host + ":9093"

func getInfoFromService1(a string) string{

	/*
	resp, err := http.PostForm("http://127.0.0.1:9091"+a,
	url.Values{"key": {"Value"}, "id": {"123"}})
	*/
	
	resp, err := http.Get("http://127.0.0.1:9091"+a)
	if err!= nil{
		//fmt.Fprintf(w,"error")
	}
	defer resp.Body.Close()

	body, err:= ioutil.ReadAll(resp.Body)
	if err != nil{
		///fmt.Fprintf(w," read body error.")
	}
	return string(body)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	body := getInfoFromService1("/string")
	fmt.Fprintf(w,string(body))
	//fmt.Fprintf(w, "Hello astaxie!") //输出到客户端的信息
}
func allHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, userService)
	//向user service请求数据
}
func historyHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, userService)
	//向user service请求数据
}
func userHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, userService)
	//向user service请求数据
}
func deviceHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, deviceService)
	//向device service请求数据
}
func main() {
	http.HandleFunc("/test",sayhelloName)
	http.HandleFunc("/all",allHandler)
	http.HandleFunc("/history", historyHandler)       
	http.HandleFunc("/user",userHandler)
	http.HandleFunc("/device",deviceHandler)
	
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}