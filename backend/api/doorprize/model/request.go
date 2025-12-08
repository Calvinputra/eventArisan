package model

import (
	"event/backend/api/doorprize/entity"
	"gitlab.universedigital.my.id/library/golang/crud/model"
)
type CreateDoorprizeRequest struct {
	BaseDoorprizeRequest
}

func (r *CreateDoorprizeRequest) ToEntity(inputter, apiType string) *entity.Doorprize {
    return &entity.Doorprize{
        AuditTrail:      model.AuditTrail{},
        EventRecid:      r.EventRecid,
        AttendanceRecid: r.AttendanceRecid,
    }
}
