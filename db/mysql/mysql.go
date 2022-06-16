package mysql

import (
	"go-todo/server/config"
	"go-todo/utl/env"
	"go-todo/utl/log"
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

	return gorm.Open(sqlDialector, &gorm.Config{
		Logger: gormLogger,
	})

}
