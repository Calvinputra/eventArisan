package repository

import (
	"event/backend/api/event/entity"
	"github.com/sirupsen/logrus"
	"event/backend/repository"
	"gorm.io/gorm"
    "errors"
	"strings"
	"fmt"
)
type EventRepository struct {
    Live repository.BaseRepository[entity.Event]
	His  repository.BaseRepository[entity.EventHis]
	Log  *logrus.Logger
	DB   *gorm.DB
}

func NewEventRepository(db *gorm.DB, log *logrus.Logger) *EventRepository {
	return &EventRepository{
		Log: log,
		DB:  db,
	}
}

func (r *EventRepository) CheckRecid(db *gorm.DB, recid string) (bool, error) {
	var event *entity.Event
	if err := db.Table(entity.Event{}.TableName()).
		Where("recid = ?", recid).
		First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}
	return true, nil
}

func (r *EventRepository) FindAllByType(eventType string) (interface{}, error) {
    eventType = strings.ToUpper(eventType)

    switch eventType {
    case "LIVE":
        var event []entity.Event
        result := r.DB.Find(&event)
        return event, result.Error

    case "HIS":
        var event []entity.EventHis
        result := r.DB.Find(&event)
        return event, result.Error

    default:
        return nil, fmt.Errorf("invalid user type: %s", eventType)
    }
}
