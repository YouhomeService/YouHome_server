package service

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os/exec"
)

func LoadRouters() {
	http.HandleFunc("/v1/voice/wxupload", wxuploadHandler)
}

func wxuploadHandler(w http.ResponseWriter,req *http.Request) {
	var temp = "YouHome_server/services/Voice/" + "decoder"
	//需要将接收到的文件放到temp指定的文件夹，并命名为sample.silk


	//调用bash，使得sample.silk转化成sample.wav
	cmd := exec.Command("sh", "../decoder/converter.sh sample.slk wav")


	//调用xunfei.go中的Xunfeiapi()函数，得到一个识别的字符串
	result := Xunfeiapi()
}