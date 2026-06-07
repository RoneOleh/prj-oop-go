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

type EventController struct {
	eventService app.EventService
}

func NewEventController(es app.EventService) EventController {
	return EventController{
		eventService: es,
	}
}

func (c EventController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ev, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		dev := r.Context().Value(DeviceKey).(domain.Device)
		room := r.Context().Value(RoomKey).(domain.Room)
		ev.DeviceId = dev.Id
		ev.RoomId = room.Id

		ev, err = c.eventService.Save(ev)
		if err != nil {
			log.Printf("EventController.Save(c.eventService.Save): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.EventDto{}.DomainToDto(ev))
	}
}

func (c EventController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		ev := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		Success(w, resources.EventDto{}.DomainToDto(ev))
	}
}

func (c EventController) FindByDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev := r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		events, err := c.eventService.FindByDeviceId(dev.Id)
		if err != nil {
			log.Printf("EventController.FindByDevice(c.eventService.FindByDeviceId): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.DomainToDtoCollection(events))
	}
}

func (c EventController) FindByPeriod() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		period := r.URL.Query().Get("period")

		var roomId *uint64
		if room, ok := r.Context().Value(RoomKey).(domain.Room); ok {
			roomId = &room.Id
		}

	energy, err := c.eventService.FindByPeriod(roomId, org.Id, period)
	if err != nil {
		log.Printf("EventController.GetEnergyByPeriod: %s", err)
		BadRequest(w, err)
		return
	}

	response := map[string]interface{}{
		"energy_kwh": energy,
		"period":     period,
	}
	Success(w, response)
		Success(w, resources.EventDto{}.DomainToDtoCollection(events))
	}
}

func (c EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		ev := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newEv, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController.Update(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		ev.Action = newEv.Action

		ev, err = c.eventService.Update(ev)
		if err != nil {
			log.Printf("EventController.Update(c.eventService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.EventDto{}.DomainToDto(ev))
	}
}

func (c EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		ev := r.Context().Value(EventKey).(domain.Event)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		err := c.eventService.Delete(ev.Id)
		if err != nil {
			log.Printf("EventController.Delete(c.eventService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
