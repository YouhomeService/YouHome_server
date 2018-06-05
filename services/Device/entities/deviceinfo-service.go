package entities

import "fmt"

func GetDevicesByRoomId(id string) [][3]string {
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

func AddDevice(name string, eid string, rid string) (error, string) {
	err := Insert(name, eid, rid)
	fmt.Println("insertout")
	deviceid := QueryDeviceId(name, eid, rid)
	return err, deviceid
}

func SetDeviceName(id string, name string) error {
	err := UpdateNameById(id, name)
	return err
}
