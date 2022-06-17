package mysql

import (
	"go-todo/internal/env"
	"go-todo/internal/log"
	"go-todo/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Connect(cfg *config.Configuration) (*gorm.DB, error) {
	connectionString := env.GetString("DB_SQL_URL")
	sqlDialector := mysql.Open(connectionString)

	gormLogger := logger.New(
		log.NewGormLogger(),
		logger.Config{
			SlowThreshold: time.Millisecond * time.Duration(cfg.Database.SlowThreshold),
		},
	)

	log.Logger.Infof("Connecting to DB at %s", connectionString)
	return gorm.Open(sqlDialector, &gorm.Config{
		Logger: gormLogger,
	})

}
