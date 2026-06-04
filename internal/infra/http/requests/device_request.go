package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	OrganizationId   uint64  `json:"orzanization_id"`
	RoomId           *uint64 `json:"room_id"`
	// Description      *string `json:"description"`
	GUID             string  `json:"guid"     validate:"required"`
	InventotyNumber  string  `json:"invemtory_number"`
	SerialNumber     string  `json:"serial_number"`
	Charachteristics string  `json:"charactescs"`
	Units            string  `json:"units"`
	Power            float64 `json:"power"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		OrganizationId:   r.OrganizationId,
		RoomId:           r.RoomId,
		// Description:      r.Description,
		GUID:             r.GUID,
		InventotyNumber:  r.InventotyNumber,
		SerialNumber:     r.SerialNumber,
		Charachteristics: r.Charachteristics,
		Units:            r.Units,
		Power:            r.Power,
	}, nil
}
