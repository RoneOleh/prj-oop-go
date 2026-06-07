package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
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
func (d EventDto) DomainToDto(e domain.Event) EventDto {
	return EventDto{
		Id:          e.Id,
		DeviceId:    e.DeviceId,
		RoomId:      e.RoomId,
		Action:      e.Action,
		CreatedDate: e.CreatedDate,
		UpdatedDate: e.UpdatedDate,
		DeletedDate: e.DeletedDate,
	}
}

func (d EventDto) DomainToDtoCollection(events []domain.Event) []EventDto {
	result := make([]EventDto, len(events))
	for i, e := range events {
		result[i] = d.DomainToDto(e)
	}
	return result
}