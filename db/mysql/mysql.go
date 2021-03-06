package mysql

import (
	"time"

	"go-todo/internal/env"
	"go-todo/internal/log"
	"go-todo/server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Configuration) (*gorm.DB, error) {
	connectionString := env.GetString("DB_SQL_URL")
	sqlDialector := mysql.Open(connectionString)

	gormLogger := logger.New(
		log.NewGormLogger(),
		logger.Config{
			SlowThreshold: time.Millisecond * time.Duration(cfg.Database.SlowThreshold),
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	log.Logger.Infof("Connecting to DB at %s", connectionString)
	return gorm.Open(sqlDialector, &gorm.Config{
		Logger: gormLogger,
	})
}
