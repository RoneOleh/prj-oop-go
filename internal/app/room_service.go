package app

import "github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"

type roomService struct {
	roomRepo database.RoomRepository
}
type RoomService interface {
	Save(room  domian.Room)(domain.Room,eror)
    Update(room  domian.Room)(domain.Room,eror)
	Delete(id uint64) error
}

type NewRoomService (roomRepo database.RoomRepository)RoomService{
    rm, err := s.orgRepo.Save(room)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Organization{}, err
	}

	return rm, nil
}
func (s roomService)Save(room  domian.Room)(domain.Room,eror){
    rm, err := s.orgRepo.Save(room)
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
func (s roomService) Delete(id uint64) error {
	err := s.orgRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}	
}
