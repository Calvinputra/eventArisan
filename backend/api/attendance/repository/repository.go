package repository

import (
	"event/backend/api/attendance/entity"
	"github.com/sirupsen/logrus"
	"event/backend/repository"
	"gorm.io/gorm"
    "errors"
	"strings"
	"fmt"
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

func (r *AttendanceRepository) FindAllByType(attendanceType string) (interface{}, error) {
    attendanceType = strings.ToUpper(attendanceType)

    switch attendanceType {
    case "LIVE":
        var attendance []entity.Attendance
        result := r.DB.Find(&attendance)
        return attendance, result.Error
    default:
        return nil, fmt.Errorf("invalid user type: %s", attendanceType)
    }
}	

func (r *AttendanceRepository) FindByIdAndType(db *gorm.DB, recid string, typeAttendance string) (*entity.Attendance, error) {
    var attendance entity.Attendance
    result := db.Table(entity.Attendance{}.TableName()).
        Where("recid = ?", recid).
        Where("record_status = ?", typeAttendance).
        First(&attendance)

    if result.Error != nil {
        return nil, result.Error
    }

    return &attendance, nil
}