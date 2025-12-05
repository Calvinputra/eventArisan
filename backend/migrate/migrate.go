package migrate

import (
	EventEntity "event/backend/api/event/entity"
	AttendanceEntity "event/backend/api/attendance/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&EventEntity.Event{},
		&EventEntity.EventHis{},
		&AttendanceEntity.Attendance{},
		&AttendanceEntity.AttendanceHis{},
	)
}
