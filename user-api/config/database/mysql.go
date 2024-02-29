package database

import (
	"github.com/OVillas/user-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MysqlConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	return db, err
}
