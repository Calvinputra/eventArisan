package model

import (
	"event/backend/api/event/entity"

	"gitlab.universedigital.my.id/library/golang/crud/model"
)

/*
GALLERY CMS RESPONSE
*/

type EventResponse struct {
	model.AuditTrail
	Name             string                `json:"Name"`
	Description      string                `json:"Description"`
	StartDateTime    int64                 `json:"StartDateTime"`
	EndDateTime      int64                 `json:"EndDateTime"`
	Status           string                `json:"Status"`
}

func (EventResponse) ToResponse(entityEvent *entity.Event) EventResponse {
	return EventResponse{
		AuditTrail:    entityEvent.AuditTrail,
		Name:           entityEvent.Name,
		Description:          entityEvent.Description,
		StartDateTime:       entityEvent.StartDateTime,
		EndDateTime:         entityEvent.EndDateTime,
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