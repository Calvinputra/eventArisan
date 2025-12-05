package config

import "github.com/sirupsen/logrus"

func NewLog() *logrus.Logger {
	log := logrus.New()
	//log.SetLevel(logrus.Level(config.GetInt("LOG_LEVEL")))
	// log.SetFormatter(&logrus.JSONFormatter{})
	return log
}
