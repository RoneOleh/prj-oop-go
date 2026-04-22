package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

type organization struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:""`
	Name        string     `db:""`
	Description string     `db:""`
	City        string     `db:""`
	Address     string     `db:""`
	Lat         float64    `db:""`
	LON         float64    `db:""`
	CreatedDate time.Time  `db:""`
	UpdatedDate time.Time  `db:""`
	DeletedDate *time.Time `db:""`
}

type organizationRepository struct {
	coll db.Collection
	sess db.Session
}

type OrganizationRepository interface {
	Save(o domain.Organization) (domain.Organization, error)
}

func NewOrganizationRepository(session db.Session) OrganizationRepository {
	return organizationRepository{
		sess: session,
		coll: session.Collection("organizations"),
	}
}

func (r organizationRepository) Save(o domain.Organization) (domain.Organization, error) {

	org := r.mapDomainToModel(o)
	now := time.Now()
	org.CreatedDate = now
	org.UpdatedDate = now

	err := r.coll.InsertReturning(&org)
	if err != nil {
		return domain.Organization{}, err
	}

	o = r.mapModelToDomain(org)
	return o, nil
}

func (r organizationRepository) mapDomainToModel(o domain.Organization) organization {
	return organization{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Address:     o.Address,
		Lat:         o.Lat,
		LON:         o.LON,
		CreatedDate: o.CreatedDate,
		UpdatedDate: o.UpdatedDate,
		DeletedDate: o.DeletedDate,
	}
}

func (r organizationRepository) mapModelToDomain(o organization) domain.Organization {
	return domain.Organization{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Address:     o.Address,
		Lat:         o.Lat,
		LON:         o.LON,
		CreatedDate: o.CreatedDate,
		UpdatedDate: o.UpdatedDate,
		DeletedDate: o.DeletedDate,
	}
}
