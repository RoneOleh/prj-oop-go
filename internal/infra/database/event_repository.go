package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const EventsTableName = "events"

type event struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      uint64     `db:"room_id"`
	Action      string     `db:"action"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type EventRepository interface {
	Save(e domain.Event) (domain.Event, error)
	Find(id uint64) (domain.Event, error)
	FindByDeviceId(deviceId uint64) ([]domain.Event, error)
	FindByRoomId(roomId uint64) ([]domain.Event, error)
	FindByPeriod(roomId *uint64, orgId uint64, from time.Time, to time.Time) ([]domain.Event, error)
	Update(e domain.Event) (domain.Event, error)
	Delete(id uint64) error
}

type eventRepository struct {
	coll db.Collection
}

func NewEventRepository(dbSession db.Session) EventRepository {
	return eventRepository{
		coll: dbSession.Collection(EventsTableName),
	}
}

func (r eventRepository) Save(e domain.Event) (domain.Event, error) {
	ev := r.mapDomainToModel(e)
	ev.CreatedDate = time.Now()
	ev.UpdatedDate = time.Now()
	err := r.coll.InsertReturning(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}

func (r eventRepository) Find(id uint64) (domain.Event, error) {
	var ev event
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}
func (r eventRepository) FindByPeriod(roomId *uint64, orgId uint64, from time.Time, to time.Time) ([]domain.Event, error) {
	var data []event

	cond := db.Cond{
		"created_date >=": from,
		"created_date <=": to,
		"deleted_date":    nil,
	}

	if roomId != nil {
		
		cond["room_id"] = *roomId
	} else {
		
		cond["room_id IN"] = db.Raw("SELECT id FROM rooms WHERE organization_id = ?", orgId)
	}


	err := r.coll.Find(cond).OrderBy("device_id", "created_date").All(&data)
	if err != nil {
		return nil, err
	}

	
	var result []domain.Event
	for _, e := range data {
		result = append(result, r.mapModelToDomain(e))
	}
	
	return result, nil
}
func (r eventRepository) FindByDeviceId(deviceId uint64) ([]domain.Event, error) {
	var data []event
	err := r.coll.Find(db.Cond{"device_id": deviceId, "deleted_date": nil}).
		OrderBy("-created_date").All(&data)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Event, len(data))
	for i, ev := range data {
		result[i] = r.mapModelToDomain(ev)
	}
	return result, nil
}

func (r eventRepository) FindByRoomId(roomId uint64) ([]domain.Event, error) {
	var data []event
	err := r.coll.Find(db.Cond{"room_id": roomId, "deleted_date": nil}).
		OrderBy("-created_date").All(&data)
	if err != nil {
		return nil, err
	}
	result := make([]domain.Event, len(data))
	for i, ev := range data {
		result[i] = r.mapModelToDomain(ev)
	}
	return result, nil
}

func (r eventRepository) Update(e domain.Event) (domain.Event, error) {
	ev := r.mapDomainToModel(e)
	ev.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": ev.Id, "deleted_date": nil}).Update(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}

func (r eventRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).
		Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r eventRepository) mapDomainToModel(e domain.Event) event {
	return event{
		Id:          e.Id,
		DeviceId:    e.DeviceId,
		RoomId:      e.RoomId,
		Action:      e.Action,
		CreatedDate: e.CreatedDate,
		UpdatedDate: e.UpdatedDate,
		DeletedDate: e.DeletedDate,
	}
}

func (r eventRepository) mapModelToDomain(e event) domain.Event {
	return domain.Event{
		Id:          e.Id,
		DeviceId:    e.DeviceId,
		RoomId:      e.RoomId,
		Action:      e.Action,
		CreatedDate: e.CreatedDate,
		UpdatedDate: e.UpdatedDate,
		DeletedDate: e.DeletedDate,
	}
}