package entities

import "encoding/json"

func AddRoom(roomName string, userId string) error {
	sql := "INSERT INTO room (roomName, userId) VALUES (?, ?)"
	_, err := mydb.Exec(sql, roomName, userId)
	return err
}

func GetRooms(Userid string) string {
	sql := "SELECT roomId,roomName FROM room WHERE userId = ?"
	rows, err := mydb.Query(sql, Userid)
	checkErr(err)

	type Fat struct {
		RoomId string `json:"roomId"`
		RoomName string `json:"roomName"`
	}
	var temp Fat
	result := make([]Fat, 0)
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&temp.RoomId,&temp.RoomName)
		checkErr(err)
		result = append(result,temp)
	}
	re , err := json.Marshal(result)
	checkErr(err)
	return string(re)
}

func GetRoomId(roomName string, userId string) string{
	sql := "SELECT roomId FROM room WHERE roomName = ? and userId = ? ORDER BY roomId"
	rows, err := mydb.Query(sql, roomName,userId)
	checkErr(err)
	var roomId string
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&roomId)
	}
	return roomId
}
func GetRoomName(roomId string)string{
	sql := "select roomName from room where roomId = ?"
	rows, err := mydb.Query(sql, roomId)
	checkErr(err)
	var roomName string
	if rows == nil{
		return ""
	}
	for rows.Next(){
		rows.Scan(&roomName)
	}
	return roomName
}
func UpdateRoomName(roomName string,roomId string)string{
	sql := "UPDATE room SET roomName = ? WHERE roomId = ?"
	_ ,err := mydb.Exec(sql,roomName,roomId)
	checkErr(err)
	return roomName
}
func DeleteRoom(id string) error {
	sql := "delete FROM room where roomId=? "
	_, err := mydb.Exec(sql, id)
	return err
}