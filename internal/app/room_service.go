package app

import "github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"

type roomService struct {
	roomRepo database.RoomRepository
}
type RoomService interface {
}

type NewRoomService (roomRepo database.RoomRepository)RoomService{
    rm, err := s.orgRepo.Save(o)
	if err != nil {
		log.Printf("organizationService.Save(s.orgRepo.Save): %s", err)
		return domain.Organization{}, err
	}

	return rm, nil
}
func (s roomService)Save(room  domian.Room)(domain.Room,eror){
    rm, err := s.orgRepo.Save(o)
	if err != nil {
		log.Printf("organizationService.Save(s.orgRepo.Save): %s", err)
		return domain.Organization{}, err
	}

	return rm, nil
}
