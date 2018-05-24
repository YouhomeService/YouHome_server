package service

import "net/http"

func LoadRoute() {
	http.HandleFunc("/v1/scenes/name",nameHandler)
	http.HandleFunc("/v1/scenes/delete",deleteHandler)
	http.HandleFunc("/v1/scenes", handler)
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	if r.Method == "GET"{
		getScenes(w,r)
	} else if r.Method == "POST"{
		addScenes(w,r)
	} else if r.Method =="DELETE"{
		deleteScenes(w,r)
	}
	return
}