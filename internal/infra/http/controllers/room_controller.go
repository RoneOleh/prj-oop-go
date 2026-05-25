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

type RoomController struct {
	roomService app.RoomService
}

func NewRoomController(rs app.RoomService) RoomController {
	return RoomController{
		roomService: rs,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rm, err := requests.Bind(r,
			requests.RoomRequest{},
			domain.Room{} )
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}
	 	
		org := r.Context().Value(OrgKey).(domain.Organization)

		rm.OrgId = org.Id
		

		rm, err = c.roomService.Save(rm)
		if err != nil {
			log.Printf("OrganizationController.Save(c.orgService.Save): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.RoomDto{}.DomainToDto(rm))
	}


    var rmDto resources.RoomDto
    rmDto = rmDto.DomainToDto(savedRoom)
	if err != nil {
		log.Printf("RoomController.FindList(c.rmService.FindList): %s", err)
		InternalServerError(w, err)
		return
	}

	Success(w, resources.RoomDto{}.DomainToDtoCollection(rms))
  }





func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)
		room := r.Context().Value(RoomKey).(domain.Room)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

     Success(w, resources.RoomDto{}.DomainToDto(room))
	}
 }




func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org := r.Context().Value(OrgKey).(domain.Organization)

		if user.Id != room.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newRoom, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		rm.Name = newRoom.Name
		rm.Description = newRoom.Description

		rm, err = c.roomService.Update(rm)
		if err != nil {
			log.Printf("RoomController.Update(c.rmService.Update): %s", err)
			BadRequest(w, err)
			return
		}

		Success(w, resources.RoomDto{}.DomainToDto(rm))
	}
}
