package entity

import (
    attendanceEntity "event/backend/api/attendance/entity"
    "gitlab.universedigital.my.id/library/golang/crud/model"
)

type Doorprize struct {
    model.AuditTrail

    EventRecid      string `gorm:"column:event_recid;type:varchar(191);not null"`
    AttendanceRecid string `gorm:"column:attendance_recid;type:varchar(191);not null;index"`

    Attendance      attendanceEntity.Attendance `gorm:"foreignKey:AttendanceRecid;references:Recid"`
}

type DoorprizeHis struct {
    model.AuditTrail

    EventRecid      string `gorm:"column:event_recid;type:varchar(191);not null"`
    AttendanceRecid string `gorm:"column:attendance_recid;type:varchar(191);not null"`
}

func (Doorprize) TableName() string {
    return "doorprize"
}

func (DoorprizeHis) TableName() string {
    return "doorprize_his"
}
