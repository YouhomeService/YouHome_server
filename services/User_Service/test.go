package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"fmt"
	"github.com/Suenaa/agenda-golang/service/tools"
	"io/ioutil"
	"encoding/json"
	"bytes"
)
func main() {

		data := struct {
			UserId string `json:"userId"`
			UserName string `json:"userName"`
		}{"abcd","lizq"}
		buf, _ := json.Marshal(data)
		res, err := http.Post("http://127.0.0.1:8080/v1/users/userName",
			"application/json", bytes.NewBuffer(buf))

		res1, err1 := http.Get("http://127.0.0.1:8080/v1/users?userId=abc")

		if err == nil {
			fmt.Println("Success")
			defer res.Body.Close()
		} else {
			tools.Report(err)
		}

		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		if err1 == nil {
			fmt.Println("Success")
			defer res.Body.Close()
		} else {
			tools.Report(err)
		}

		body1, _ := ioutil.ReadAll(res1.Body)
		fmt.Println(string(body1))


}


