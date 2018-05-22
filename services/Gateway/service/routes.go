package service

import (
	"net/http"
)

func InitRoutes(){
	http.HandleFunc("/scene/all",sceneAllHandler)
	http.HandleFunc("/scene/create",sceneCreateHandler)
	http.HandleFunc("/scene/delete",deleteSceneHandler)
	http.HandleFunc("/scene/name",sceneNameHandler)

	http.HandleFunc("/device/all",deviceAllHandler)
	http.HandleFunc("/device",deviceHandler)


	http.HandleFunc("/user/login",userLoginHandler)
	http.HandleFunc("/user",userHandler)
}

func checkError(err error) {
	if err != nil{
		panic(err)
	}
}