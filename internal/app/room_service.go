package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type roomService struct {
	roomRepo database.RoomRepository
}
type RoomService interface {
	Save(room  domain.Room)(domain.Room,eror)
    Update(room domain.Room)(domain.Room,eror)
	Delete(id uint64) error
}

func NewRoomService(rm database.RoomRepository) RoomService  {
    return RoomService{
		roomRepo:rm,
	}
}
func (s roomService)Save(room  domian.Room)(domain.Room,error){
    rm, err := roomRepo.Save(room)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
}
func (s roomService)Update(room  domian.Room)(domain.Room,eror){
    rm, err := s.orgRepo.Update(room)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
func (s roomService) Find(id uint64) (interface{}, error) {
	rm, err := s.roomRepo.Find(id)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.Find): %s", err)
		return nil, err
	}
   rm.Rooms,err = s.roomRepo.FindByOrgId(org.Id)
	if err != nil{
		log.Printf("roomService.Find(s.roomRepo.FindByOrg): %s",err)
	return nil,err
	}

	return rm, nil
}
func (s roomService) FindList(uId uint64) ([]domain.Organization, error) {
	rms, err := s.rmRepo.FindList(uId)
	if err != nil {
		log.Printf("roomService.FindList(s.rmRepo.FindList): %s", err)
		return nil, err
	}

	return rms, nil
}


func (s roomService) Delete(id uint64) error {
	err := s.orgRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}	
}
