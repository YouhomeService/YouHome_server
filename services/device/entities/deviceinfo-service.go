package entities

func GetDevicesBySceneId(id string) [][3]string {
	device := Query(id)
	return device
}

func GetEntityId(id string) string {
	return QueryByDeviceId(id)
}
