package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementTableName = "measurements"

type measurement struct {
	Id             uint64         `db:"id,omitempty"`
	RoomId         *uint64        `db:"room_id"`
	DeviceId       uint64         `db:"device_id"`
	Value          float64        `db:"value"`
	Type           string         `db:"type"`
	CreatedDate    time.Time      `db:"created_date,omitempty"`
	UpdatedDate    time.Time      `db:"updated_date,omitempty"`
	DeletedDate    *time.Time     `db:"deleted_date,omitempty"`
	
}

type MeasurementRepository interface {
	FindByDeviceId(deviceId uint64) ([]domain.Measurement, error)
	FindByDeviceAndFromTime(deviceId uint64,from time.Time) ([]domain.Measurement, error)
	Save(m domain.Measurment) (domain.Measurement, error)
	Delete(id uint64) error
	Update(m domain.Measurement) (domain.Measurement, error) 
}

type measurementRepository struct {
	coll db.Collection
	
}

func NewMeasurementRepository(dbsession db.Session) MeasurementRepository {
	return measurmentRepository{
		coll: session.Collection(MeasurmentTableName),
		
	}
}
func (r measurementRepository) FindByDeviceId(deviceId uint64) ([]domain.Measurement, error) {
	var data []measurment

	err := r.coll.
		Find(db.Cond{
			"device_id": deviceId,
			"deleted_date": nil,
           }).
		 OrderBy("-created_date").All(&data)
	     if err != nil {
		return nil, err
	}

	result := make([]domain.Measurement, len(data))
	for i,meas:=range data{
		result [i] = r.mapModelToDomain(meas)
	}
	return result, nil
}

func (r measurementRepository) FindByDeviceIdAndFromTime(deviceId uint64,from time.Time)([]domain.Measurement, error)  {
	var data [] measurment

	err := r.coll.
		Find(db.Cond{
			"device_id": deviceId,
			"created_date >=" : from,
			"deleted_date": nil, }).OrderBy("-created_date")
		  .All(&data)

	if err != nil {
		return nil, err
	}
	result := make([]domain.Measurement, len(data))
	for i,meas:= range data{
		result [i] = r.mapModelToDomain(meas)
	}
	return result, nil
}
func (r measurementRepository ) Save(m domain.Measurement ) ([]domain.Measurement, error) {
	meas:= r.mapDomainToModel(m)
	meas.CreatedDate:= r.time.Now(m)
	

	err := r.coll.InsertReturning(&meas)
	if err != nil {
		return domain.Measurement{}, err
	}

	d = r.mapModelToDomain(meas)
	return d, nil
}
func (r *measurementRepository) Delete(id uint64) error {
	now := time.Now()

	return r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil,
		}).
		Update(map[string]interface{}{
			"deleted_date": now,
		})
}

func (r measurementRepository ) Update(m domain.Measurement ) ([]domain.Measurement, error) {
	meas:= r.mapDomainToModel(m)
	meas.UpdatedDate:= r.time.Now()
	Find(db.Cond{
			"id": measId,
           }).
		   Update(meas)
	     if err != nil {
		return nil, err
	}
  if err != nil {
		return domain.Measurement{}, err
	}

	return  r.mapModelToDomain(meas), nil
}


func (r measurementRepository) mapDomainToModel(m domain.Measurement) measurement {
	return measurement{
	Id:                m.Id ,                         
    DeviceId:          m.DeviceId,              
	RoomId:      	   m.RoomId,                 
	Value:       	   m.Value,              
	Type:       	   m.Type ,
	CreatedDate:       m.CreatedDate,
	UpdatedDate:       m.UpdatedDate,
	DeletedDate:       m.DeletedDate,
    
  }
		
}
func (r measurementRepository) mapModelToDomain(m measurement)  domain.Measurement {
	return  domain.Measurement{
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


