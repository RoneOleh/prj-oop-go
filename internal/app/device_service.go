package app

import (
	"log"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type deviceService struct {
	deviceRepo database.DeviceRepository
}
type DeviceService interface {
	Save(d domain.Device)(domain.Device,error)
    Update(d domain.Device)(domain.Device,error)
	Delete(id uint64) error
	Find(id uint64) (interface{}, error) 

}

func NewDeviceService(dev database.DeviceRepository) DeviceService  {
    return &deviceService {
		deviceRepo: dev, 
	}
}
func (s deviceService)Save(dev  domain.Device)(domain.Device,error){
    dev, err := s.deviceRepo.Save(dev)
	if err != nil {
		log.Printf("deviceService.Save(s.deviceRepo.Save): %s", err)
		return domain.Device{}, err
	}

	return dev, nil
}
func (s deviceService)Update(dev domain.Device)(domain.Device,error){
    dev, err := s.deviceRepo.Update(dev)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Device{}, err
	}

	return dev, nil
}

func(s deviceService) Find(id uint64) (interface{}, error) {
	dev, err := s.deviceRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.deviceRepo.Find): %s", err)
		return nil, err
	}
  
	return dev, nil
}


func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("deviceService.Delete(s.deviceRepo.Delete): %s", err)
		return err
	}	
	return err
}
