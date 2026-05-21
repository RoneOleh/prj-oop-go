package domain

type Device struct {
	Id             uint64
	OrganizationId uint64
	RoomId         *uint64
	GUID            string
	InventotyNumber string
	SerialNumber    string
	Charachteristics string
	Category         DeviceCatogory
	Units
    Power

}

type DeviceCatogory string


const{
	Sensor DevivceCategory = SENSOR
	Actuator DeviceCatogoty = ACTUATOR
	
}