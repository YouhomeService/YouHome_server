package entities


func GetDevicesByRoomId(id string) [][4]string {
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

func AddDevice(name string, eid string, rid string, url string) (error, string) {
	err := Insert(name, eid, rid, url)
	deviceid := QueryDeviceId(name, eid, rid)
	return err, deviceid
}

func DeleteDevice(id string) error {
	err := Delete(id)
	return err
}

func SetDeviceName(id string, name string) error {
	err := UpdateNameById(id, name)
	return err
}
