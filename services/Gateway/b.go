package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func f(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数, 默认是不会解析的
	fmt.Println(r.Form) //这些是服务器端的打印信息
	fmt.Println("path", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}



	
	if r.Method == "GET"{
		fmt.Fprintf(w, "this is device service.Method: Get") //输出到客户端的信息
	}else{
		fmt.Fprintf(w, "this is device service.Method: POST") //输出到客户端的信息
	}
	re, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, string(re))}

func main() {
	http.HandleFunc("/", f) 
	
	err := http.ListenAndServe(":9092", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}