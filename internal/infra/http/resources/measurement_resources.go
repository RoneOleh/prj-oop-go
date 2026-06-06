package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type MeasurementDto struct {
    Id                 uint64                  `json:"id"`
    DeviceId           uint64                  `json:"device_id"`
	RoomId            *uint64                  `json:"room_id"`
	Value              float64                 `json:"value"` 
	Type               string                  `json:"type"`
	CreatedDate        time.Time               `json:"created_date"`
	UpdatedDate        time.Time               `json:"updated_date"`  
	DeletedDate        *time.Time              `json:"deleted_date"` 
}
func (d MeasurementDto ) DomainToDto( m domain.Measurement) MeasurementDto {
	return MeasurementDto{
    Id:                m.Id ,                         
    DeviceId:          m.DeviceId,              
	RoomId:      	   m.RoomId,                 
	Value:       	   m.Value,              
	Type:       	   m.Type,
	CreatedDate:       m.CreatedDate,
    UpdatedDate:       m.UpdatedDate,
	DeletedDate:       m.DeletedDate,         
 }
}


func (d MeasurementDto) DomainToDtoCollection(measurments []domain.Measurement) []MeasurementDto {
	result:= make([]MeasurementDto, len(measurments))
    for i, m := range measurments  {
		result[i] = d.DomainToDto(m)
	}

	return result
}

