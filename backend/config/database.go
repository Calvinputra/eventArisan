package config

import (
	"event/backend/helper"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FJakarta&charset=utf8mb4&timeout=30s",
		helper.GetStringEnv("DB_USERNAME"),
		helper.GetStringEnv("DB_PASSWORD"),
		helper.GetStringEnv("DB_HOST"),
		helper.GetStringEnv("DB_PORT"),
		helper.GetStringEnv("DB_NAME"),
	)

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB from gorm.DB: " + err.Error())
	}

	// optional: pooling config
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
