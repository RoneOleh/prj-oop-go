package resources

import (
	"time"

	
)

type EventDto struct {
	Id                 uint64           `json:"id"`
	DeviceId           uint64           `json:"device_id"`
	RoomId             uint64           `json:"room_id"`
	Action             string           `json:"action"`
	CreatedDate        time.Time        `json:"created_date"`
	UpdatedDate        time.Time        `json:"updated_date"`  
	DeletedDate        *time.Time       `json:"deleted_date"` 

}