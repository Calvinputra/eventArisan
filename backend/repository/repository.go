package repository

import (
	"event/backend/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) List(db *gorm.DB, entity *[]T) error {
	return db.Find(entity).Error
}

func (r *BaseRepository[T]) ListBasedOnPartnerRecid(db *gorm.DB, entity *[]T, partnerRecid string) error {

	switch partnerRecid {
	case constants.PartnerRecidAll:
		return db.Find(entity).Error
	}

	return db.Where("partner_recid = ?", partnerRecid).Find(entity).Error
}

func (r *BaseRepository[T]) ListBasedOnPartnerRecidWithPUBLIC(db *gorm.DB, entity *[]T, partnerRecid string) error {

	switch partnerRecid {
	case constants.PartnerRecidAll:
		return db.Find(entity).Error
	case constants.PartnerRecidPublic:
		return db.Where("partner_recid = ?", constants.PartnerRecidPublic).Find(entity).Error
	}

	return db.Where("partner_recid = ?", partnerRecid).Find(entity).Error
}

func (r *BaseRepository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(db *gorm.DB, entity *T, recid string) error {
	return db.Where("recid = ?", recid).Save(entity).Error
}

func (r *BaseRepository[T]) Delete(db *gorm.DB, entity *T, recid string) error {
	return db.Where("recid = ?", recid).Delete(entity).Error
}

func (r *BaseRepository[T]) FindById(db *gorm.DB, entity *T, id string) error {
	return db.Where("recid = ?", id).Find(entity).Error
}

func (r *BaseRepository[T]) FindByPartnerRecid(db *gorm.DB, entity *[]T, partnerRecid string) error {
	return db.Where("partner_recid = ?", partnerRecid).Find(entity).Error
}

func (r *BaseRepository[T]) GetAll(db *gorm.DB) ([]*T, error) {
	var entities []*T
	if err := db.Find(&entities).Error; err != nil {
		panic(err)
	}
	return entities, nil
}

func (r *BaseRepository[T]) BulkCreate(db *gorm.DB, entities []T) error {
	return db.Create(&entities).Error
}

func (r *BaseRepository[T]) BulkUpsert(db *gorm.DB, entities []T, columns []string) error {
	// Use reflection to get the column names from the struct type	// Perform bulk upsert
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "recid"}},
		DoUpdates: clause.AssignmentColumns(columns),
	}).Create(&entities).Error
}

func (r *BaseRepository[T]) BulkDelete(db *gorm.DB, entities []T) error {
	return db.Delete(&entities).Error
}
