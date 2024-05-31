package storage

import (
	"context"
	"errors"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type IDbmStorage interface {
	ListUsers(ctx context.Context, offset int64, limit int) ([]model.MUserModel, error)
}

type dbmStorage struct {
	configDb config.DatabaseConfig
	dbm      *gorm.DB
}

func NewDbmStorage(cnf config.DatabaseConfig, dbm *gorm.DB) IDbmStorage {
	return dbmStorage{
		configDb: cnf,
		dbm:      dbm,
	}
}

func (d dbmStorage) table() *gorm.DB {
	return d.dbm.Table(fmt.Sprintf("%s.%s", d.configDb.Schema, model.MUserModel{}.TableName()))
}

// ListUsers implements IDbmStorage.
func (d dbmStorage) ListUsers(ctx context.Context, offset int64, limit int) ([]model.MUserModel, error) {
	var result []model.MUserModel
	err := d.table().Select("user_id", "loyalty_id", "email", "phone", "cur_original_id").Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}
	return result, nil
}
