package entities

import "encoding/json"

func AddScene(sceneName string, userId string) error {
	sql := "INSERT INTO scene (sceneName, userId) VALUES (?, ?)"
	_, err := mydb.Exec(sql, sceneName, userId)
	return err
}

func GetScenes(Userid string) string {
	sql := "SELECT sceneId,sceneName FROM scene WHERE userId = ?"
	rows, err := mydb.Query(sql, Userid)
	checkErr(err)

	type Fat struct {
		SceneId string `json:"sceneId"`
		SceneName string `json:"sceneName"`
	}
	var temp Fat
	result := make([]Fat, 0)
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&temp.SceneId,&temp.SceneName)
		checkErr(err)
		result = append(result,temp)
	}
	re , err := json.Marshal(result)
	checkErr(err)
	return string(re)
}

func GetSceneId(sceneName string, userId string) string{
	sql := "SELECT sceneId FROM scene WHERE sceneName = ? and userId = ? ORDER BY sceneId"
	rows, err := mydb.Query(sql, sceneName,userId)
	checkErr(err)
	var sceneId string
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&sceneId)
	}
	return sceneId
}
func GetSceneName(sceneId string)string{
	sql := "select sceneName from scene where sceneId = ?"
	rows, err := mydb.Query(sql, sceneId)
	checkErr(err)
	var sceneName string
	if rows == nil{
		return ""
	}
	for rows.Next(){
		rows.Scan(&sceneName)
	}
	return sceneName
}
func UpdateSceneName(sceneName string,sceneId string)string{
	sql := "UPDATE scene SET sceneName = ? WHERE sceneId = ?"
	_ ,err := mydb.Exec(sql,sceneName,sceneId)
	checkErr(err)
	return sceneName
}
func DeleteScene(id string) error {
	sql := "delete FROM scene where sceneId=? "
	_, err := mydb.Exec(sql, id)
	return err
}