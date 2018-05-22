package entities



func Login(userId string) error {

	var err error

	if GetNameById(userId) == "" {
		err = AddUser(userId, userId)
	} else {
		err = nil
	}

	return err
}

func GetUserName(id string) string {

	name := GetNameById(id)

	return name
}

func SetUserName(id string, name string) error {

	err := UpdateNameById(id, name)

	return err
}