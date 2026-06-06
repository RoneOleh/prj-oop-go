package domain

import "time"


type Device struct {
	Id               uint64
	OrganizationId   uint64
	RoomId           *uint64
	GUID             string
	InventotyNumber  string
	SerialNumber     string
	Charachteristics string
	Category         DeviceCategory
	Units            string
	Power            float64
	CreatedDate      time.Time
	UpdatedDate      time.Time
	DeletedDate      *time.Time
}

type DeviceCategory string

const (
	Sensor   DeviceCategory = "SENSOR"
	Actuator DeviceCategory  = "ACTUATOR"
)
