package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitDB(dsn string, lgr logger.Writer) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			lgr,
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		SkipDefaultTransaction: true, // Epic performance improvement
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Albi{}, &OldPrice{}, &SearchQuery{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
