package repository

import (
	"event/backend/api/doorprize/entity"
	"github.com/sirupsen/logrus"
	"event/backend/repository"
	"gorm.io/gorm"
    "errors"
	"strings"
	"fmt"
)
type DoorprizeRepository struct {
    Live repository.BaseRepository[entity.Doorprize]
	Log  *logrus.Logger
	DB   *gorm.DB
}

func NewDoorprizeRepository(db *gorm.DB, log *logrus.Logger) *DoorprizeRepository {
	return &DoorprizeRepository{
		Log: log,
		DB:  db,
	}
}

func (r *DoorprizeRepository) CheckRecid(db *gorm.DB, recid string) (bool, error) {
	var doorprize *entity.Doorprize
	if err := db.Table(entity.Doorprize{}.TableName()).
		Where("recid = ?", recid).
		First(&doorprize).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}
	return true, nil
}

func (r *DoorprizeRepository) FindAllByType(doorprizeType string) (interface{}, error) {
    doorprizeType = strings.ToUpper(doorprizeType)

    switch doorprizeType {
    case "LIVE":
        var doorprize []entity.Doorprize
        result := r.DB.Find(&doorprize)
        return doorprize, result.Error
    default:
        return nil, fmt.Errorf("invalid user type: %s", doorprizeType)
    }
}

func (r *DoorprizeRepository) FindAllByEvent(eventRecid string) ([]entity.Doorprize, error) {
    var result []entity.Doorprize
    err := r.DB.Where("event_recid = ?", eventRecid).Find(&result).Error
    return result, err
}

