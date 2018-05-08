package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
	"bytes"
	"encoding/json"
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
		
		re, _ := ioutil.ReadAll(r.Body)
		temp := bytes.NewBuffer(re).String()
		fmt.Fprintf(w, "this is user service.Method: Get."+temp) //输出到客户端的信息
	}else{
		re, _ := ioutil.ReadAll(r.Body)
		
		fmt.Fprintf(w, "this is user service.Method: POST."+string(re)) //输出到客户端的信息
	}

	data := struct{
		UserId string `json:"userId"`
		UserName string `json:"userName"`
	}{"15331181","liangfeng"}
	buf ,_:= json.Marshal(data)
	fmt.Fprintln(w,string(buf))
	
}

func main() {
	http.HandleFunc("/", f) 
	
	err := http.ListenAndServe(":9091", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}