package service

import (
	"net/http"
	"fmt"
	"os/exec"
    "encoding/json" 
    "regexp"
	"os"
	"io"
)

type Product struct {
    Code      string  `json:"code"`
    Data string   `json:"data,string"`
    Number    int     `json:"number,string"`
    Price     float64 `json:"price,string"`
    IsOnSale  bool    `json:"is_on_sale,string"`
}

func LoadRouters() {
	http.HandleFunc("/v1/voice/wxupload", wxuploadHandler)
}

func wxuploadHandler(w http.ResponseWriter,req *http.Request) {
	var temp = "YouHome_server/services/Voice/" + "decoder"
	//需要将接收到的文件放到temp指定的文件夹，并命名为sample.silk
		req.ParseMultipartForm(32 << 20)
        file, handler, err := req.FormFile("uploadfile")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer file.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        f, err := os.OpenFile(temp+"sample.slk", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()
        io.Copy(f, file)


	//调用bash，使得sample.silk转化成sample.wav
	exec.Command("sh", "../decoder/converter.sh sample.slk wav")


	//调用xunfei.go中的Xunfeiapi()函数，得到一个识别的字符串
	result := Xunfeiapi()

	fmt.Println(result)

    actiton := regexp.MustCompile(`(?:开|关)`)
    actitonresult := actiton.FindAllString(result, -1)

    room := regexp.MustCompile(`(?:卧室|卫生间|客厅|阳台|书房|主卧|次卧)`)
    roomresult := room.FindAllString(result, -1)

    device := regexp.MustCompile(`(?:灯|空调|智能网关|风扇)`)
    deviceresult := device.FindAllString(result, -1)
}