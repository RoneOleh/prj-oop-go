package app
/*
import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.mcom/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type EventService interface {
	Save(e domain.Event) (domain.Event, error)
	Find(id uint64) (domain.Event, error)
	//FindByDeviceId(deviceId uint64) ([]domain.Event, error)
	// FindByPeriod(roomId *uint64, orgId uint64, period string) (float64, error)
	Update(e domain.Event) (domain.Event, error)
	Delete(id uint64) error
}

type eventService struct {
	eventRepo  database.EventRepository
	deviceRepo database.DeviceRepository
}

func NewEventService(er database.EventRepository, dr database.DeviceRepository) EventService {
	return eventService{
		eventRepositor: er,
		deviceRepositor: dr,
	}
}

func (s eventService) Save(e domain.Event) (domain.Event, error) {
	ev, err := s.eventRepo.Save(e)
	if err != nil {
		log.Printf("eventService.Save: %s", err)
		return domain.Event{}, err
	}
	return ev, nil
}

func (s eventService) Find(id uint64) (domain.Event, error) {
	ev, err := s.eventRepo.Find(id)
	if err != nil {
		log.Printf("eventService.Find: %s", err)
		return domain.Event{}, err
	}
	return ev, nil
}

// func (s eventService) FindByDeviceId(deviceId uint64) ([]domain.Event, error) {
// 	events, err := s.eventRepo.FindByDeviceId(deviceId)
// 	if err != nil {
// 		log.Printf("eventService.FindByDeviceId: %s", err)
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (s eventService) FindByPeriod(
// 	roomId *uint64,
// 	orgId uint64,
// 	period string,
// ) (float64, error) {

// 	now := time.Now()

// 	var from time.Time
// 	var to time.Time

// 	switch period {
// 	case "day":
// 		from = time.Date(
// 			now.Year(),
// 			now.Month(),
// 			now.Day(),
// 			0, 0, 0, 0,
// 			now.Location(),
// 		)
// 		to = from.AddDate(0, 0, 1)

// 	case "week":
// 		weekday := int(now.Weekday())
// 		if weekday == 0 {
// 			weekday = 7
// 		}

// 		from = time.Date(
// 			now.Year(),
// 			now.Month(),
// 			now.Day()-weekday+1,
// 			0, 0, 0, 0,
// 			now.Location(),
// 		)
// 		to = from.AddDate(0, 0, 7)

// 	case "month":
// 		from = time.Date(
// 			now.Year(),
// 			now.Month(),
// 			1,
// 			0, 0, 0, 0,
// 			now.Location(),
// 		)
// 		to = from.AddDate(0, 1, 0)

// 	default:
// 		return 0, errors.New(
// 			"invalid period: use 'day', 'week' or 'month'",
// 		)
// 	}

// 	devices, err := s.deviceRepo.FindByDeviceId(
// 		orgId,
// 		roomId,
// 	)
// 	if err != nil {
// 		log.Printf("eventService.GetEnergyByPeriod: %s", err)
// 		return 0, err
// 	}

// 	var totalEnergy float64

// 	for _, device := range devices {
// 		energy, err := s.eventRepo.FindByPeriod(
// 			device.Id,
// 			from,
// 			to,
// 		)
// 		if err != nil {
// 			log.Printf("eventService.FindByPeriod: %s", err)
// 			return 0, err
// 		}

// 		totalEnergy += energy
// 	}

// 	return totalEnergy, nil
// }

func (s eventService) Update(e domain.Event) (domain.Event, error) {
	ev, err := s.eventRepo.Update(e)
	if err != nil {
		log.Printf("eventService.Update: %s", err)
		return domain.Event{}, err
	}
	return ev, nil
}

func (s eventService) Delete(id uint64) error {
	err := s.eventRepo.Delete(id)
	if err != nil {
		log.Printf("eventService.Delete: %s", err)
		return err
	}
	return nil
}
*/