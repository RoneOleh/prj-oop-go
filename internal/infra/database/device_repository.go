package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id               uint64  `db:"id,omitempty"`
	OrganizationId   uint64  `db:"organization_id"`
	RoomId           *uint64 `db:"room_id"`
	GUID             string  `db:"guid"`
	InventotyNumber  string  `db:"IN"`
	SerialNumber     string
	Charachteristics string
	Category         DeviceCatogory
	Units            string
	Power            float64
	CreatedDate      time.Time
	UpdatedDate      time.Time
	DeletedDate      *time.Time
}

type RoomRepository interface {
	FindByOrgId(oId uint64) ([]domain.Room, error)
	Save(o domain.Room) (domain.Room, error)
	Update(o domain.Room) (domain.Room, error)
	Delete(id uint64) error
	Find(oId uint64) (domain.Room, error)
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(session db.Session) RoomRepository {
	return roomRepository{
		coll: session.Collection(RoomsTableName),
		sess: session,
	}
}
func (r roomRepository) Find(oId uint64) (domain.Room, error) {
	var rm room

	err := r.coll.
		Find(db.Cond{"id": oId, "deleted_date": nil}).
		One(&rm)
	if err != nil {
		return domain.Room{}, err
	}

	o := r.mapModelToDomain(rm)
	return o, nil
}

func (r roomRepository) FindByOrgId(oId uint64) ([]domain.Room, error) {
	var rooms []room

	err := r.coll.
		Find(db.Cond{
			"organization_id": oId,
			"deleted_date":    nil,
		}).All(&rooms)

	if err != nil {
		return nil, err
	}
	rms := r.mapModelToDomainCollection(rooms)
	return rms, nil
}
func (r roomRepository) Save(o domain.Room) (domain.Room, error) {
	rm := r.mapDomainToModel(o)
	now := time.Now()
	rm.CreatedDate = now
	rm.UpdatedDate = now

	err := r.coll.InsertReturning(&rm)
	if err != nil {
		return domain.Room{}, err
	}

	o = r.mapModelToDomain(rm)
	return o, nil
}
func (r roomRepository) Update(o domain.Room) (domain.Room, error) {
	rm := r.mapDomainToModel(o)
	rm.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&rm)
	if err != nil {
		return domain.Room{}, err
	}
	o = r.mapModelToDomain(rm)
	return o, nil
}
func (r roomRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(rm room) domain.Room {
	return domain.Room{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
		CreatedDate:    rm.CreatedDate,
		UpdatedDate:    rm.UpdatedDate,
		DeletedDate:    rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	rms := make([]domain.Room, len(rooms))
	for i := range rooms {
		rms[i] = r.mapModelToDomain(rooms[i])
	}
	return rms
}
