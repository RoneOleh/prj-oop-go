package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type device  struct {
	Id               uint64               `db:"id,omitempty"`
	OrganizationId   uint64               `db:"organization_id"`
	RoomId           *uint64              `db:"room_id"`
	GUID             string               `db:"guid"`
	InventotyNumber  string               `db:"inventoty_numder"`
	SerialNumber     string              `db:"serial_numder"`
	Charachteristics string              `db:"charachteristics"` 
	Category         domain.DeviceCategory     `db:"category"` 
	Units            string               `db:"units"` 
	Power            float64              `db:"power"` 
	CreatedDate      time.Time            `db:"created_date"` 
	UpdatedDate      time.Time            `db:"update_date"` 
	DeletedDate      *time.Time            `db:"deleted_date"`
}

type DeviceRepository interface {
	FindByOrgId(oId uint64) ([]domain.Device, error)
	Save(o domain.Device) (domain.Device, error)
	Update(o domain.Device) (domain.Device, error)
	Delete(id uint64) error
	Find(oId uint64) (domain.Device, error)
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(session db.Session) DeviceRepository {
	return deviceRepository{
		coll: session.Collection(DevicesTableName),
		sess: session,
	}
}
func (r deviceRepository) Find(oId uint64) (domain.Device, error) {
	var dev device 

	err := r.coll.
		Find(db.Cond{"id": oId, "deleted_date": nil}).
		One(&dev)
	if err != nil {
		return domain.Device{}, err
	}

	d := r.mapModelToDomain(dev)
	return d, nil
}

func (r deviceRepository) FindByOrgId(oId uint64) ([]domain.Device, error) {
	var devices []device 

	err := r.coll.
		Find(db.Cond{
			"organization_id": oId,
			"deleted_date":    nil,
		}).All(&devices)

	if err != nil {
		return nil, err
	}
	devs := r.mapModelToDomainCollection(devices)
	return devs, nil
}
func (r deviceRepository) Save(d domain.Device) (domain.Device, error) {
	dev:= r.mapDomainToModel(d)
	now := time.Now()
	dev.CreatedDate = now
	dev.UpdatedDate = now

	err := r.coll.InsertReturning(&dev)
	if err != nil {
		return domain.Device{}, err
	}

	d = r.mapModelToDomain(dev)
	return d, nil
}
func (r deviceRepository) Update(d domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(d)
	dev.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": d.Id, "deleted_date": nil}).Update(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	 d = r.mapModelToDomain(dev)
	return d, nil
}
func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device {
    Id:                  d.Id,               
	OrganizationId:   	d.OrganizationId,   
	RoomId:           	d.RoomId,           
	GUID:            	d.GUID,            
	InventotyNumber:  	d.InventotyNumber,  
	SerialNumber:     	d.SerialNumber,     
	Charachteristics: 	d.Charachteristics, 
	Category:        	d.Category,        
	Units:         	    d.Units,         
	Power:        	    d.Power,        
	CreatedDate:      	d.CreatedDate,      
	UpdatedDate:      	d.UpdatedDate,      
	DeletedDate:      	d.DeletedDate,   
  }
		
}


func (r deviceRepository) mapModelToDomain(dev device ) domain.Device {
	return domain.Device{
	Id:                 dev.Id,               
	OrganizationId:   	dev.OrganizationId,   
	RoomId:           	dev.RoomId,           
	GUID:            	dev.GUID,            
	InventotyNumber:  	dev.InventotyNumber,  
	SerialNumber:     	dev.SerialNumber,     
	Charachteristics: 	dev.Charachteristics, 
	Category:        	dev.Category,        
	Units:         	    dev.Units,         
	Power:        	    dev.Power,        
	CreatedDate:      	dev.CreatedDate,      
	UpdatedDate:      	dev.UpdatedDate,      
	DeletedDate:      	dev.DeletedDate,    
	}
}

func (r deviceRepository) mapModelToDomainCollection(devices []device ) []domain.Device {
	devs := make([]domain.Device, len(devices))
	for i := range devices {
		devs[i] = r.mapModelToDomain(devices[i])
	}
	return devs
}
