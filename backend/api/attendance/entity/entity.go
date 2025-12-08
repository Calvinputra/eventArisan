package entity

import "gitlab.universedigital.my.id/library/golang/crud/model"

type Attendance struct {
	model.AuditTrail
	EventRecid        string `gorm:"column:event_recid;not null"`
	Name              string `gorm:"column:name;not null"`
	Code              string `gorm:"column:code;not null;unique"`
	NoTable          int64  `gorm:"column:no_table;not null"`
	StatusCheckin     int8 `gorm:"column:status_checkin;not null" json:"statusCheckin"`
	StatusSouvenir   int8 `gorm:"column:status_souvenir;not null" json:"statusSouvenir"`
	CheckinTime      int64  `gorm:"column:checkin_time;not null"`
	SouvenirTime      int64  `gorm:"column:souvenir_time;not null"`
}

type AttendanceHis struct {
	model.AuditTrail
	EventRecid        string `gorm:"column:event_recid;not null"`
	Name              string `gorm:"column:name;not null"`
	Code              string `gorm:"column:code;not null"`
	NoTable          int64  `gorm:"column:no_table;not null"`
	StatusCheckin     int8 `gorm:"column:status_checkin;not null"`
	StatusSouvenir   int8 `gorm:"column:status_souvenir;not null"`
	CheckinTime      int64  `gorm:"column:checkin_time;not null"`
	SouvenirTime      int64  `gorm:"column:souvenir_time;not null"`
}

func (Attendance) TableName() string {
	return "attendance"
}

func (AttendanceHis) TableName() string {
	return "attendance_his"
}