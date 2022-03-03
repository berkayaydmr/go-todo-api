package common

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(DatabaseUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(DatabaseUrl), &gorm.Config{})

	logger := zap.NewExample()

	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return db
}
