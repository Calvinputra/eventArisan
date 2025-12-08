package model

import (
	"event/backend/api/doorprize/entity"
	"gitlab.universedigital.my.id/library/golang/crud/model"
)
type DoorprizeResponse struct {
    model.AuditTrail
    EventRecid       string `json:"eventRecid"`
    AttendanceRecid  string `json:"attendanceRecid"`
    AttendanceCode   string `json:"attendanceCode"`
    AttendanceName   string `json:"attendanceName"`
}

func (DoorprizeResponse) ToResponse(e *entity.Doorprize) DoorprizeResponse {
    return DoorprizeResponse{
        AuditTrail:      e.AuditTrail,
        EventRecid:      e.EventRecid,
        AttendanceRecid: e.AttendanceRecid,
        AttendanceCode:  e.Attendance.Code,
        AttendanceName:  e.Attendance.Name,
    }
}


func (DoorprizeResponse) ToResponseList(doorprizeList []entity.Doorprize) []DoorprizeResponse {
	responseList := []DoorprizeResponse{}
	for i := range doorprizeList {
		responseList = append(responseList, DoorprizeResponse{}.ToResponse(&doorprizeList[i]))
	}
	return responseList
}
