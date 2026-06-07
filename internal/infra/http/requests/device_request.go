package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	OrganizationId   uint64  `json:"organization_id"`
	RoomId           *uint64 `json:"room_id"`
	// Description      *string `json:"description"`
	GUID             string  `json:"guid"     validate:"required"`
	InventotyNumber  string  `json:"inventory_number"`
	SerialNumber     string  `json:"serial_number"`
	Charachteristics string  `json:"charachteristics"`
	Category         string  `json:"category"`
	Units            string  `json:"units"`
	Power            float64 `json:"power"`
}

func (d DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		OrganizationId:   d.OrganizationId,
		RoomId:           d.RoomId,
		// Description:   d  r.Description,
		GUID:             d.GUID,
		InventotyNumber:  d.InventotyNumber,
		SerialNumber:     d.SerialNumber,
		Charachteristics: d.Charachteristics,
		Category:         domain.DeviceCategory(d.Category),
		Units:            d.Units,
		Power:            d.Power,
	}, nil
}
