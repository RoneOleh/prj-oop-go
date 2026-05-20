package app

import (
	"log"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type roomService struct {
	roomRepo database.RoomRepository
}
type RoomService interface {
	Save(room  domain.Room)(domain.Room,error)
    Update(room domain.Room)(domain.Room,error)
	Delete(id uint64) error
	Find(id uint64) (interface{}, error) 
	FindList(uId uint64) ([]domain.Room, error)
}

func NewRoomService(rm database.RoomRepository) RoomService  {
    return roomService{
		roomRepo:rm,
	}
}
func (s roomService)Save(room  domain.Room)(domain.Room,error){
    rm, err := s.roomRepo.Save(room)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
}
func (s roomService)Update(room domain.Room)(domain.Room,error){
    rm, err := s.roomRepo.Update(room)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
}

func(s roomService) Find(id uint64) (interface{}, error) {
	rm, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return nil, err
	}
  
	return rm, nil
}
func (s roomService) FindList(uId uint64) ([]domain.Room, error) {
	rms, err := s.roomRepo.FindList(uId)
	if err != nil {
		log.Printf("roomService.FindList(s.roomRepo.FindList): %s", err)
		return nil, err
	}

	return rms, nil
}


func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}	
}
