package storage

import (
	"context"
	"errors"
	"fmt"
	"golang-server/config"
	"golang-server/module/api/dto"
	"golang-server/pkg/logger"

	"gorm.io/gorm"
)

type CommonStorageParams struct {
	TableName    string
	Query        *gorm.DB
	CommonFilter dto.CommonFilter
	Filter       interface{}
	Data         interface{}
}

type commonStorage struct {
	db       *gorm.DB
	configDb config.DatabaseConfig
}

func (s commonStorage) table(tableName string) *gorm.DB {
	return s.db.Table(fmt.Sprintf("%s.%s", s.configDb.Schema, tableName))
}

func (s commonStorage) CFindOne(ctx context.Context, param CommonStorageParams) error {
	logger.Info(ctx, fmt.Sprintf("CFindOne %s table", param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
	if len(param.CommonFilter.Select) > 0 {
		param.Query = param.Query.Select(param.CommonFilter.Select)
	}
	tx := param.Query.First(param.Data)
	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("find one %s error", param.TableName))
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return tx.Error
}

func (s commonStorage) CList(ctx context.Context, param CommonStorageParams) error {
	logger.Info(ctx, fmt.Sprintf("CList %s table", param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
	if param.CommonFilter.Limit != 0 {
		param.Query = param.Query.Limit(param.CommonFilter.Limit)
	}
	if param.CommonFilter.Offset != nil {
		param.Query = param.Query.Offset(*param.CommonFilter.Offset)
	}
	if param.CommonFilter.Sort != "" {
		param.Query = param.Query.Order(param.CommonFilter.Sort)
	}
	if len(param.CommonFilter.Select) > 0 {
		param.Query = param.Query.Select(param.CommonFilter.Select)
	}
	tx := param.Query.Find(param.Data) // day la contro vao bien ket qua
	if tx.Error != nil {
		logger.Error(
			ctx,
			tx.Error,
			fmt.Sprintf("list %s error", param.TableName),
		)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return tx.Error
}

func (s commonStorage) CInsert(ctx context.Context, param CommonStorageParams) error {
	logger.Info(ctx, fmt.Sprintf("CInsert %s table", param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
	tx := s.table(param.TableName).Create(param.Data)
	if tx.Error != nil {
		logger.Error(
			ctx,
			tx.Error,
			fmt.Sprintf("insert %s data error", param.TableName),
		)
	}
	return tx.Error
}

func (s commonStorage) CUpdateMany(ctx context.Context, param CommonStorageParams) error {
	logger.Info(ctx, fmt.Sprintf("CUpdateMany %s table", param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
	tx := param.Query.Updates(param.Data)

	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("Failed to update many %s", param.TableName))
	}
	return tx.Error
}

func (s commonStorage) CDelete(ctx context.Context, param CommonStorageParams) error {
	logger.Info(ctx, fmt.Sprintf("CDelete %s table", param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
	tx := param.Query.Delete(param.Data)
	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("Failed to delete %s", param.TableName))
	}
	return tx.Error
}
