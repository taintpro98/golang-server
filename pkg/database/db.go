package database

import (
	"fmt"
	"golang-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresqlDatabase(databaseCnf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := GetDatabaseDSN(databaseCnf)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func GetDatabaseDSN(cnf config.DatabaseConfig) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cnf.Host,
		cnf.Username,
		cnf.Password,
		cnf.DatabaseName,
		cnf.Port,
	)
	return dsn
}
