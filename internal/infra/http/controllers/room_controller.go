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
			domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		org := r.Context().Value(OrgKey).(domain.Organization)

		rm.OrganizationId = org.Id

		rm, err = c.roomService.Save(rm)
		if err != nil {
			log.Printf("OrganizationController.Save(c.orgService.Save): %s", err)
			InternalServerError(w, err)
			return
		}
		Success(w, resources.RoomDto{}.DomainToDto(rm))
	}

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
		room := r.Context().Value(RoomKey).(domain.Room)

		if user.Id != org.UserId {
			Forbidden(w, errors.New("access denied"))
			return
		}
		if org.Id != room.OrganizationId {
			Forbidden(w, errors.New("access denied"))
			return
		}

		newRoom, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Update(requests.Bind): %s", err)
			BadRequest(w, err)
			return
		}

		room.Name = newRoom.Name
		room.Description = newRoom.Description


		room, err = c.roomService.Update(room)
		if err != nil {
			log.Printf("RoomController.Update(c.rmService.Update): %s", err)
			BadRequest(w, err)
			return
		}

		Success(w, resources.RoomDto{}.DomainToDto(room))
	}
}

	func (c RoomController) Delete() http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
      user := r.Context().Value(UserKey).(domain.User)
      org := r.Context().Value(OrgKey).(domain.Organization)
      room:= r.Context().Value(RoomKey).(domain.Room)
  

    if user.Id != org.UserId {
      Forbidden(w, errors.New("access denied"))
      return
    }
	if org.Id != room.OrganizationId {
    Forbidden(w, errors.New("access denied"))
    return

}

     err := c.roomService.Delete(room.Id)
      if err != nil {
      InternalServerError(w, err)
      return
    }

    noContent(w)
  }
}