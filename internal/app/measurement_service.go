package app

import (
	"errors"
	"log"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type measurementService struct {
	measurementRepo database.MeasurementRepository
}

type MeasurementService interface {
	FindByDeviceId(deviceId uint64) ([]domain.Measurement, error)
	GetHistoryByPeriod(deviceId uint64, period string) ([]domain.Measurement, error)
	Save(m domain.Measurement) (domain.Measurement, error)
	Update(m domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
}

func NewMeasurementService(mr database.MeasurementRepository) MeasurementService {
	return measurementService{
		measurementRepo: mr,
	}
}

func (s measurementService) FindByDeviceId(deviceId uint64) ([]domain.Measurement, error) {
	measurements, err := s.measurementRepo.FindByDeviceId(deviceId)
	if err != nil {
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) GetHistoryByPeriod(deviceId uint64, period string) ([]domain.Measurement, error) {
	now := time.Now()
	var (
		measurements []domain.Measurement
		err          error
	)

	switch period {
	case "day":
		measurements, err = s.findByDeviceAndDay(deviceId, now)
	case "week":
		measurements, err = s.findByDeviceAndWeek(deviceId, now)
	case "month":
		measurements, err = s.findByDeviceAndMonth(deviceId, now)
	default:
		return nil, errors.New("invalid period: use 'day', 'week' or 'month'")
	}

	if err != nil {
		log.Printf("MeasurementService.GetHistoryByPeriod(%s): %v", period, err)
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) findByDeviceAndDay(deviceId uint64, date time.Time) ([]domain.Measurement, error) {
	from := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	to := from.Add(24 * time.Hour)
	return s.measurementRepo.FindByDeviceAndPeriod(deviceId, from, to)
}

func (s measurementService) findByDeviceAndWeek(deviceId uint64, date time.Time) ([]domain.Measurement, error) {
	weekday := int(date.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	from := time.Date(date.Year(), date.Month(), date.Day()-weekday+1, 0, 0, 0, 0, date.Location())
	to := from.Add(7 * 24 * time.Hour)
	return s.measurementRepo.FindByDeviceAndPeriod (deviceId, from, to)
}

func (s measurementService) findByDeviceAndMonth(deviceId uint64, date time.Time) ([]domain.Measurement, error) {
	from := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	to := from.AddDate(0, 1, 0)
	return s.measurementRepo.FindByDeviceAndPeriod (deviceId, from, to)
}

func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	meas, err := s.measurementRepo.Save(m)
	if err != nil {
		return domain.Measurement{}, err
	}
	return meas, nil
}

func (s measurementService) Update(m domain.Measurement) (domain.Measurement, error) {
	meas, err := s.measurementRepo.Update(m)
	if err != nil {
		return domain.Measurement{}, err
	}
	return meas, nil
}

func (s measurementService) Delete(id uint64) error {
	err := s.measurementRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}