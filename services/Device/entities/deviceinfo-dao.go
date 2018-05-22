package entities

func Query(id string) [][3]string{
	sql := "SELECT deviceid, devicename, entityid FROM deviceinfo WHERE sceneid = ?"
	rows, err := mydb.Query(sql, id)
	checkErr(err)
	var dList [][3]string
	if rows == nil {
		return dList
	}
	for rows.Next() {
		var temp [3]string
		rows.Scan(&temp[0], &temp[1], &temp[2])
		dList = append(dList, temp)
	}
	return dList
}

func Insert(devicename string, entityid string, sceneid string) error {
	sql := "INSERT INTO deviceinfo (devicename, entityid, sceneid) VALUES (?, ?, ?)"
	_, err := mydb.Exec(sql, devicename, entityid, sceneid)
	return err
}

func QueryByDeviceId(id string) (string, string) {
	sql := "SELECT entityid, devicename FROM deviceinfo WHERE deviceid = ?"
	rows, err := mydb.Query(sql, id)
	checkErr(err)
	var entityId, deviceName string
	if rows == nil {
		return "",""
	}
	for rows.Next() {
		rows.Scan(&entityId, &deviceName)
	}
	return entityId, deviceName
}

func UpdateNameById(id string, name string) error {
	sql := "UPDATE deviceinfo set devicename=? where deviceid=? "
	_, err := mydb.Exec(sql, name, id)
	return err
}