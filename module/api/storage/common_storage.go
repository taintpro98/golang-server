package storage

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/pkg/logger"

	"gorm.io/gorm"
)

type commonStorage struct {
	db       *gorm.DB
	configDb config.DatabaseConfig
}

func (s commonStorage) table(tableName string) *gorm.DB {
	return s.db.Table(fmt.Sprintf("%s.%s", s.configDb.Schema, tableName))
}

func (s commonStorage) BuildQuery(filter interface{}) *gorm.DB {
	panic("")
}

func (s commonStorage) CInsert(ctx context.Context, tableName string, data interface{}) error {
	logger.Info(ctx, fmt.Sprintf("%s insert", tableName))
	err := s.table(tableName).Create(data).Error
	if err != nil {
		logger.Error(
			ctx,
			err,
			fmt.Sprintf("insert %s data error", tableName),
		)
	}
	return err
}

func (s commonStorage) CUpdateMany(ctx context.Context, tableName string, filter interface{}, data interface{}) error {
	logger.Info(ctx, fmt.Sprintf("%s insert", tableName))
	query := s.BuildQuery(filter)
	tx := query.Updates(data)

	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("Failed to update %s", tableName))
		return tx.Error
	}
	return nil
}
