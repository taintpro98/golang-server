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

func (s commonStorage) log(ctx context.Context, funcName string, param CommonStorageParams) {
	logger.Info(ctx, fmt.Sprintf("%s %s table", funcName, param.TableName), logger.LogField{
		Key:   "data",
		Value: param.Data,
	}, logger.LogField{
		Key:   "filter",
		Value: param.Filter,
	}, logger.LogField{
		Key:   "common_filter",
		Value: param.CommonFilter,
	})
}

func (s commonStorage) CCount(ctx context.Context, param CommonStorageParams) (*int64, error) {
	s.log(ctx, "CCount", param)
	var count int64
	tx := param.Query.Count(&count)
	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("count %s error", param.TableName))
		return nil, tx.Error
	}
	return &count, nil
}

func (s commonStorage) CFindOne(ctx context.Context, param CommonStorageParams) error {
	s.log(ctx, "CFindOne", param)
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
	s.log(ctx, "CList", param)
	if param.CommonFilter.Limit != 0 {
		param.Query = param.Query.Limit(param.CommonFilter.Limit)
	}
	if param.CommonFilter.Offset != nil {
		param.Query = param.Query.Offset(*param.CommonFilter.Offset)
	}
	if param.CommonFilter.Sort != "" {
		param.Query = param.Query.Order(param.CommonFilter.Sort) // age desc hoac age asc hoac age
	}
	if len(param.CommonFilter.Select) > 0 {
		param.Query = param.Query.Select(param.CommonFilter.Select)
	}
	tx := param.Query.Find(param.Data) // day la con tro vao bien ket qua
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
	s.log(ctx, "CInsert", param)

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
	s.log(ctx, "CUpdateMany", param)

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
