package entities

func AddUser(userId string, userName string) error {
	sql := "INSERT INTO userinfo (userid, username) VALUES (?, ?)"
	_, err := mydb.Exec(sql, userId, userName)
	return err
}

func GetNameById(id string) string {
	sql := "SELECT username FROM userinfo WHERE userid = ?"
	rows, err := mydb.Query(sql, id)
	checkErr(err)
	var username string
	if rows == nil {
		return ""
	}
	for rows.Next() {
		rows.Scan(&username)
	}
	return username
}

func UpdateNameById(id string, name string) error {
	sql := "UPDATE userinfo set username=? where userid=? "
	_, err := mydb.Exec(sql, name, id)
	return err
}