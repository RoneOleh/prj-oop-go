package app

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type measurementService struct {
	measurementRepo database.MeasurementRepository
}
type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	Update(m domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
	FindByDeviceId(deviceId uint64) ([]domain.Measurement, error)
	FindByDeviceAndFromTime(deviceId uint64, from time.Time) ([]domain.Measurement, error)
}

func NewMeasurementService(m database.MeasurementRepository) DeviceService {
	return measurementService{
		measRepo: mr,
	}
}
func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	meas, err := s.measurementRepo.Save(m)
	if err != nil {
		return domain.Measurement{}, err
	}
	return meas, nil
}

func (s measurementService) FindByDeviceId(deviceId uint64) ([]domain.Measurement, error) {
	measurements, err := s.measurementRepo.FindByDeviceId(deviceId)
	if err != nil {
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) FindByDeviceAndFromTime(deviceId uint64, from time.Time) ([]domain.Measurement, error) {
	measurements, err := s.measurementRepo.FindByDeviceAndFromTime(deviceId, from)
	if err != nil {
		return nil, err
	}
	return measurements, nil
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
