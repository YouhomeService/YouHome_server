package main
import (
	"net/http"
	"io/ioutil"
	"github.com/tidwall/gjson"
	"strings"
	"encoding/json"
)

func main() {
	resp, _ := http.Get("http://localhost:8123/api/states")
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	jsonStream := string(body)

	var re []string
	result := gjson.Get(jsonStream, "#.entity_id")
	for _, name := range result.Array() {
		//println(name.String())
		id := name.String()
		temp :=strings.Split(id,".")
		if temp[0] == "switch" || temp[0] == "light" || temp[0]=="sensor"{
			//println(name.String())
			re = append(re, name.String())
		}
	}
	data, _ :=json.Marshal(re)
	println(string(data))
}