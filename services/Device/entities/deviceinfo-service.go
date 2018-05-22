package entities

func GetDevicesBySceneId(id string) [][3]string {
	device := Query(id)
	return device
}

func GetEntityId(id string) string {
	entityId, _ := QueryByDeviceId(id)
	return entityId
}

func GetDeviceName(id string) string {
	_, deviceName := QueryByDeviceId(id)
	return deviceName
}

func SetDeviceName(id string, name string) error {
	err := UpdateNameById(id, name)
	return err
}
