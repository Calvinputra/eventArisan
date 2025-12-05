package model

import (
	"event/backend/api/event/entity"
	"gitlab.universedigital.my.id/library/golang/crud/model"
)
type CreateEventRequest struct {
	BaseEventRequest
}

func (r *CreateEventRequest) ToEntity(inputter, apiType string) *entity.Event {
	return &entity.Event{
		AuditTrail:  model.AuditTrail{},
		Name:        r.Name,
		Description: r.Description,
		StartDateTime:  r.StartDateTime,
		EndDateTime:    r.EndDateTime,
		Status:         r.Status,
	}
}	

type UpdateEventRequest struct {
	Recid string `json:"recid" validate:"required"`
	BaseEventRequest
}

func (r *UpdateEventRequest) ToEntity(currentEvent *entity.Event, apiType string) *entity.Event {
	currentEvent.Name = r.Name
	currentEvent.Description = r.Description
	currentEvent.StartDateTime = r.StartDateTime
	currentEvent.EndDateTime = r.EndDateTime
	currentEvent.Status = r.Status

	return currentEvent 
}



type ListEventRequest struct {
	Recid string `json:"recid" validate:"required"`
	BaseEventRequest
}

func (r *ListEventRequest) ToEntity(currentEvent *entity.Event, apiType string) *entity.Event {
	currentEvent.Name = r.Name
	currentEvent.Description = r.Description
	currentEvent.StartDateTime = r.StartDateTime
	currentEvent.EndDateTime = r.EndDateTime
	currentEvent.Status = r.Status

	return currentEvent 
}
