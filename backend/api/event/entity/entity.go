package entity

import "gitlab.universedigital.my.id/library/golang/crud/model"

type Event struct {
    model.AuditTrail
    Name           string `gorm:"column:name;not null"`
    Description    string `gorm:"column:description"`
    StartDateTime  int64  `gorm:"column:start_date;not null"`
    EndDateTime    int64  `gorm:"column:end_date;not null"`
    Status         string `gorm:"column:status;not null"`
}

type EventHis struct {
    model.AuditTrail
    Name           string `gorm:"column:name;not null"`
    Description    string `gorm:"column:description"`
    StartDateTime  int64  `gorm:"column:start_date;not null"`
    EndDateTime    int64  `gorm:"column:end_date;not null"`
    Status         string `gorm:"column:status;not null"`
}

func (Event) TableName() string {
	return "event"
}

func (EventHis) TableName() string {
	return "event_his"
}
