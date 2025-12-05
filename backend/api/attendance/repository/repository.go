package repository

import (
	"event/backend/api/attendance/entity"
	"github.com/sirupsen/logrus"
	"event/backend/repository"
	"gorm.io/gorm"
    "errors"
)
type AttendanceRepository struct {
    Live repository.BaseRepository[entity.Attendance]
	His  repository.BaseRepository[entity.AttendanceHis]
	Log  *logrus.Logger
	DB   *gorm.DB
}

func NewAttendanceRepository(log *logrus.Logger, db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{
		Log: log,
		DB: db,
	}
}

func (r *AttendanceRepository) CheckRecid(db *gorm.DB, recid string) (bool, error) {
	var attendance *entity.Attendance
	if err := db.Table(entity.Attendance{}.TableName()).
		Where("recid = ?", recid).
		First(&attendance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}
	return true, nil
}

func (r *AttendanceRepository) BulkCreate(db *gorm.DB, entities []entity.Attendance) error {
	return db.Create(&entities).Error
}

func (r *AttendanceRepository) GetByEventRecidAndCode(eventRecid, code string) (*entity.Attendance, error) {
    var attendance entity.Attendance

    err := r.DB.
        Table(entity.Attendance{}.TableName()).
        Where("event_recid = ? AND code = ?", eventRecid, code).
        First(&attendance).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }

    return &attendance, nil
}
