package storage

import (
	"context"
	"errors"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/pkg/logger"

	"github.com/rs/zerolog/log"

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

func (s commonStorage) table(ctx context.Context, tableName string) *gorm.DB {
	return s.db.Table(fmt.Sprintf("%s.%s", s.configDb.Schema, tableName)).WithContext(ctx)
}

func (s commonStorage) log(ctx context.Context, funcName string, param CommonStorageParams) {
	log.Info().
		Ctx(ctx).
		Interface("data", param.Data).
		Interface("filter", param.Filter).
		Interface("common_filter", param.CommonFilter).
		Msg(fmt.Sprintf("%s %s table", funcName, param.TableName))
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
	if len(param.CommonFilter.Preloads) > 0 {
		for _, item := range param.CommonFilter.Preloads {
			param.Query = param.Query.Preload(item)
		}
	}

	tx := param.Query.First(param.Data)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.Error(ctx, tx.Error, fmt.Sprintf("find one %s error", param.TableName))
		return tx.Error
	}
	return nil
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
	if len(param.CommonFilter.Preloads) > 0 { // khong khuyen khich dung preloads trong list
		for _, item := range param.CommonFilter.Preloads {
			param.Query = param.Query.Preload(item)
		}
	}

	tx := param.Query.Find(param.Data) // day la con tro vao bien ket qua
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		logger.Error(
			ctx,
			tx.Error,
			fmt.Sprintf("list %s error", param.TableName),
		)
		return tx.Error
	}
	return nil
}

func (s commonStorage) CInsert(ctx context.Context, param CommonStorageParams) error {
	s.log(ctx, "CInsert", param)

	tx := s.table(ctx, param.TableName).Create(param.Data)
	if tx.Error != nil {
		logger.Error(
			ctx,
			tx.Error,
			fmt.Sprintf("insert %s data error", param.TableName),
		)
	}
	return tx.Error
}

func (s commonStorage) CInsertBatch(ctx context.Context, param CommonStorageParams) error {
	s.log(ctx, "CInsertBatch", param)
	// Create a transaction
	tx := s.table(ctx, param.TableName).Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Info(ctx, fmt.Sprintf("rollback CInsertBatch %s", param.TableName))
			tx.Rollback()
		}
	}()
	err := tx.Create(param.Data).Error // param.Data phai la con tro
	if err != nil {
		logger.Error(ctx, err, fmt.Sprintf("CInsertBatch %s error", param.TableName))
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Error(ctx, err, fmt.Sprintf("CInsertBatch %s Error when commit transaction", param.TableName))
		return err
	}
	return nil
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
	s.log(ctx, "CDelete", param)
	tx := param.Query.Delete(param.Data)
	if tx.Error != nil {
		logger.Error(ctx, tx.Error, fmt.Sprintf("Failed to delete %s", param.TableName))
	}
	return tx.Error
}
