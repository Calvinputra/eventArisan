package model

import (
	"event/backend/api/attendance/entity"
	"gitlab.universedigital.my.id/library/golang/crud/model"
)
type CreateAttendanceRequest struct {
	BaseAttendanceRequest
}

func (r *CreateAttendanceRequest) ToEntity(inputter, apiType string) *entity.Attendance {
	return &entity.Attendance{
		AuditTrail:  model.AuditTrail{},
		EventRecid:        r.EventRecid,
		Name:        r.Name,
		Code:              r.Code,
		NoTable:           r.NoTable,
		StatusCheckin:     r.StatusCheckin,
		StatusSouvenir:    r.StatusSouvenir,
		CheckinTime:       r.CheckinTime,
		SouvenirTime:       r.SouvenirTime,
	}
}	

type UpdateAttendanceRequest struct {
	Recid string `json:"recid" validate:"required"`
	BaseAttendanceRequest
}

func (r *UpdateAttendanceRequest) ToEntity(currentAttendance *entity.Attendance, apiType string) *entity.Attendance {
	currentAttendance.EventRecid = r.EventRecid
	currentAttendance.Name = r.Name
	currentAttendance.Code = r.Code
	currentAttendance.NoTable = r.NoTable
	currentAttendance.StatusCheckin = r.StatusCheckin
	currentAttendance.StatusSouvenir = r.StatusSouvenir
	currentAttendance.CheckinTime = r.CheckinTime
	currentAttendance.SouvenirTime = r.SouvenirTime

	return currentAttendance 
}

type ScanAttendanceRequest struct {
   	EventRecid string `json:"event_recid" form:"eventRecid"`
	Code       string `json:"code" form:"code"`
	Souvenir   bool `json:"souvenir" form:"souvenir"`
}

func (r *ScanAttendanceRequest) ToEntity(currentAttendance *entity.Attendance, apiType string) *entity.Attendance {
	currentAttendance.EventRecid = r.EventRecid
	currentAttendance.Code = r.Code

	return currentAttendance 
}