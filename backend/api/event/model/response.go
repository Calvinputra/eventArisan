package model

import (
	"event/backend/api/event/entity"

	"gitlab.universedigital.my.id/library/golang/crud/model"
)


type EventResponse struct {
	model.AuditTrail
	Name             string                `json:"Name"`
	Description      string                `json:"Description"`
	Location      string                `json:"Location"`
	StartDateTime    int64                 `json:"StartDateTime"`
	Status           string                `json:"Status"`
}

func (EventResponse) ToResponse(entityEvent *entity.Event) EventResponse {
	return EventResponse{
		AuditTrail:    entityEvent.AuditTrail,
		Name:           entityEvent.Name,
		Description:          entityEvent.Description,
		Location:          entityEvent.Location,
		StartDateTime:       entityEvent.StartDateTime,
		Status: entityEvent.Status,
	}
}

func (EventResponse) ToResponseList(eventList []entity.Event) []EventResponse {
	responseList := []EventResponse{}
	for i := range eventList {
		responseList = append(responseList, EventResponse{}.ToResponse(&eventList[i]))
	}
	return responseList
}