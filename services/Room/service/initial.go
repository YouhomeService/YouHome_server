package service

import "net/http"

func LoadRoute() {
	http.HandleFunc("/v1/rooms/name",nameHandler)
	http.HandleFunc("/v1/rooms/delete",deleteHandler)
	http.HandleFunc("/v1/rooms", handler)
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if r.Method == "GET"{
		getRooms(w,r)
	} else if r.Method == "POST"{
		addRooms(w,r)
	} else if r.Method =="DELETE"{
		deleteRoom(w,r)
	}
	return
}