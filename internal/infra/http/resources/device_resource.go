package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DeviceDto struct {
    OrganizationId   uint64                    `json:"organization_id"`
	RoomId           *uint64                   `json:"room_id"`
	GUID             string                    `json:"guid"     validate:"required"`
	InventotyNumber  string                    `json:"inventory_number"`
	SerialNumber     string                    `json:"serial_number"`
	Charachteristics string                    `json:"charachteristics"`
	Category         domain.DeviceCategory     `json:"category"`
	Units            string                    `json:"units"`
	Power            float64                   `json:"power"`
	CreatedDate      time.Time                 `json:"created_date"`
	UpdatedDate      time.Time                 `json:"updated_date"`
}

func (d DeviceDto) DomainToDto(dev domain.Device) DeviceDto {
	return DeviceDto{
	    OrganizationId:   dev.OrganizationId,
		RoomId:           dev.RoomId,
        GUID:             dev.GUID,
		InventotyNumber:  dev.InventotyNumber,
		SerialNumber:     dev.SerialNumber,
		Charachteristics: dev.Charachteristics,
		Category:         domain.DeviceCategory(dev.Category),
		Units:            dev.Units,
		Power:            dev.Power,
		CreatedDate:       dev.CreatedDate,
        UpdatedDate:       dev.UpdatedDate,
  }
}


func (d DeviceDto) DomainToDtoCollection(devices[]domain.Device) []DeviceDto {
	result:= make([]DeviceDto, len(devices))
    for i,dev := range devices {
		result[i] = d.DomainToDto(dev)
	}

	return result
}
