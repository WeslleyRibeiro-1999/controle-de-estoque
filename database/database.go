package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(dsn string, models []interface{}) (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	database.AutoMigrate(models...)

	return database, nil
}
