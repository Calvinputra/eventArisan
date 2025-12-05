package model

import (
	"event/backend/api/attendance/entity"

	"gitlab.universedigital.my.id/library/golang/crud/model"
)

/*
GALLERY CMS RESPONSE
*/

type AttendanceResponse struct {
	model.AuditTrail
	eventRecid        string `gorm:"column:event_recid;not null"`
	Name              string `gorm:"column:name;not null"`
	Code              string `gorm:"column:code;not null"`
	NoTable          int64  `gorm:"column:no_table;not null"`
	StatusCheckin     int8 `gorm:"column:status_checkin;not null"`
	StatusSouvenir   int8 `gorm:"column:status_souvenir;not null"`	
	CheckinTime      int64  `gorm:"column:checkin_time;not null"`
}

func (AttendanceResponse) ToResponse(entityAttendance *entity.Attendance) AttendanceResponse {
	return AttendanceResponse{
		AuditTrail:    entityAttendance.AuditTrail,
		eventRecid: entityAttendance.EventRecid,
		Name: entityAttendance.Name,
		Code: entityAttendance.Code,
		NoTable: entityAttendance.NoTable,
		StatusCheckin: entityAttendance.StatusCheckin,
		StatusSouvenir: entityAttendance.StatusSouvenir,
		CheckinTime: entityAttendance.CheckinTime,
	}
}

func (AttendanceResponse) ToResponseList(attendanceList []entity.Attendance) []AttendanceResponse {
	responseList := []AttendanceResponse{}
	for i := range attendanceList {
		responseList = append(responseList, AttendanceResponse{}.ToResponse(&attendanceList[i]))
	}
	return responseList
}