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

type DeviceController struct {
	deviceService app.DeviceService
}
type DeviceController interface {
	deviceService app.DeviceService
	deviceService app.DeviceService
}
}


func NewDeviceController(rs app.DeviceService) domain.Device {
	return DeviceController{
		deviceService: rs,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rm, err := requests.Bind(r,
			requests.DeviceRequest{},
			domain.Device{} )
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		
	 	org := r.Context().Value(OrgKey).(domain.Organization)

		dev.OrganizationId = org.Id
		

		dev, err = c.deviceService.Save(dev)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceService.Save): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.RoomDto{}.DomainToDto(rm))
	}

}





func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev:= r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

     Success(w, resources.DeviceDto{}.DomainToDto(device))
	}
 }




func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		dev:= r.Context().Value(DeviceKey).(domain.Device)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}
		if org.Id != room.OrganizationId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newDev,err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		dev.RoomId = newDev.RoomId
	    dev.GUID= newDev.GUID
        dev.InventotyNumber= newDev.InventotyNumber
		dev.SerialNumber= newDev.SerialNumber
		dev.Category = newDev.Category
	    dev.Units = newDev.Units
        dev.Power = newDev.Power

		room, err = c.roomService.Update(dev)
		if err != nil {
			log.Printf("DeviceController.Update(c.deviceService.Update): %s", err)
			BadRequest(w, err)
			return
		}

		Success(w, resources.DeviceDto{}.DomainToDto(dev))
	}
}

