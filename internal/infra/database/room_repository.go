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
}
type roomRepository struct {
	coll db.Collection
	sess db.Session
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
