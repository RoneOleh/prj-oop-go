package requests


type DeviceRequest struct {
	RoomId      *uint64    `json:"room_id"`
	Description *string   `json:"description"`
	GUID        string    `json:"guid"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Room{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}