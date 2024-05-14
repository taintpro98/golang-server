package database

import (
	"context"
	"fmt"
	"golang-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlDatabase(ctx context.Context, databaseCnf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := GetDatabaseDSN(databaseCnf)
	client, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
		},
	), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
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
