package interfaces

import (
    "event/backend/api/attendance/model"
	basemodel "event/backend/model"
	"mime/multipart"
)

type AttendanceInterface interface {
	CreateAttendance(input *model.CreateAttendanceRequest, inputter string) basemodel.Response
	UpdateAttendance(input *model.UpdateAttendanceRequest, inputter string) basemodel.Response
	DeleteAttendance(recid string, inputter string) basemodel.Response
	ListAttendance(userType string) basemodel.Response
	ImportAttendance(eventRecid string, fileHeader *multipart.FileHeader, inputter string) basemodel.Response
	ScanAttendance(input *model.ScanAttendanceRequest, inputter string) basemodel.Response
	GetAttendance(userType string, recid string) basemodel.Response
}