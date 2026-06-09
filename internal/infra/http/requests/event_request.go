package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventRequest struct {
	Id                 uint64           `json:"id"`
	DeviceId           uint64           `json:"device_id"`
	RoomId             uint64           `json:"room_id"`
    Action             string           `json:"action" validate:"required"`
  
	
}
func (r EventRequest) ToDomainModel() (interface{}, error) {
	      return domain.Event{
	      Id:                r.Id ,                         
          DeviceId:          r.DeviceId,              
	      RoomId:      	     r.RoomId,                 
	      Action:       	 r.Action,              
	      

	}, nil
}
