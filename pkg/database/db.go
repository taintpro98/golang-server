package database

import (
	"fmt"
	"golang-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresqlDatabase(databaseCnf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		databaseCnf.Host,
		databaseCnf.Username,
		databaseCnf.Password,
		databaseCnf.DatabaseName,
		databaseCnf.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
