package main

import (
	"strings"
	"fmt"
)


func main() {
	path := "/v1/users?userId=1533"
	data := strings.Split(path,"/")[2]
	//service := strings.Split(data[2],"?")[0]
	fmt.Println(data)
}