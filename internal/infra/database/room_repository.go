package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type RoomsRepository interface {
	FindByOrgId(oId uint64) ([]domain.Room, error)
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(session db.Session) RoomRepository {
       return roomRepository{
		coll:session.Collection(RoomsTableName),
		sess:session,
	   }
}


func(r roomRepository)FindByOrgId(oId uint64)([]domain.Room,error){
	var rooms []room

	err := r.coll.
		Find(db.Cond{
			"organization_id": oId,
			"deleted_date":    nil,
		}).All(&rooms)

	if err != nil {
		return nil, err
	}
	rms :=r.mapModelToDomainCollection(rooms)
	return rms,nil
}

    
func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:          rm.Id,
		UserId:      rm.UserId,
		Name:        rm.Name,
		Description: rm.Description,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(rm room) domain.Room {
	return domain.Room{
		Id:          rm.Id,
		UserId:      rm.UserId,
		Name:        rm.Name,
		Description: rm.Description,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	rms:= make([]domain.Room, len(rooms))
	for i := range rooms {
		rms[i] = r.mapModelToDomain(rooms[i])
	}
	return rms
}