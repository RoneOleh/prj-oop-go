package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type MeasurementController struct {
	measurementService app.MeasurementService
}

func NewMeasurementController(ms app.MeasurementService) MeasurementController {
	return MeasurementController{
		measurementService: ms,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meas, err := requests.Bind(r,
			requests.MeasurementRequest{},
			domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		dev := r.Context().Value(DeviceKey).(domain.Device)
		meas.DeviceId = dev.Id

		meas, err = c.measurementService.Save(meas)
		if err != nil {
			log.Printf("MeasurementController.Save(c.measurementService.Save): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}

func (c MeasurementController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		meas := r.Context().Value(MeasurementKey).(domain.Measurement)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}
		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}

func (c MeasurementController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		meas := r.Context().Value(MeasurementKey).(domain.Measurement)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newMeas, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController.Update(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		meas.Value = newMeas.Value
		meas.Type = newMeas.Type

		meas, err = c.measurementService.Update(meas)
		if err != nil {
			log.Printf("MeasurementController.Update(c.measurementService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.MeasurementDto{}.DomainToDto(meas))
	}
}

func (c MeasurementController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		meas := r.Context().Value(MeasurementKey).(domain.Measurement)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}
		err := c.measurementService.Delete(meas.Id)
		if err != nil {
			log.Printf("MeasurementController.Delete(c.measurementService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}
		noContent(w)
	}
}