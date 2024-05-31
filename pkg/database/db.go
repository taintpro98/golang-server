package database

import (
	"context"
	"fmt"
	"golang-server/config"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewPostgresqlDatabase(ctx context.Context, databaseCnf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := GetDatabaseDSN(databaseCnf)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Sử dụng log.New để tạo logger mới
		logger.Config{
			SlowThreshold:             200,         // Định nghĩa thời gian tối thiểu để log câu truy vấn là câu truy vấn chậm
			LogLevel:                  logger.Info, // Đặt log level là logger.Info để ghi log câu truy vấn SQL
			IgnoreRecordNotFoundError: true,        // Bỏ qua lỗi Record Not Found
			Colorful:                  true,        // Sử dụng màu sắc cho log
		},
	)
	client, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		},
	), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", databaseCnf.Schema),
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = client.DB()
	if err != nil {
		return nil, err
	}
	return client, err
}

func GetDatabaseDSN(DBConf config.DatabaseConfig) string {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s TimeZone=%s",
		DBConf.Host, DBConf.Port, DBConf.Username, DBConf.DatabaseName, "UTC",
	)

	if DBConf.SSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%s", DBConf.SSLMode)
	}

	if DBConf.Password != "" {
		dsn += fmt.Sprintf(" password=%s", DBConf.Password)
	}
	return dsn
}
