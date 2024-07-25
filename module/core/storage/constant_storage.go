package storage

import (
	"context"
	"errors"
	"golang-server/config"
	"golang-server/module/core/model"
	"golang-server/pkg/constants"

	"gorm.io/gorm"
)

type IConstantStorage interface {
	FindOne(ctx context.Context, code constants.GolangServerConstant) (model.ConstantModel, error)

	UpdateMany(ctx context.Context, code constants.GolangServerConstant, value string) error
}

type constantStorage struct {
	commonStorage
}

func NewConstantStorage(cfg config.DatabaseConfig, db *gorm.DB) IConstantStorage {
	return constantStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

// FindOne implements IConstantStorage.
func (s constantStorage) FindOne(ctx context.Context, code constants.GolangServerConstant) (model.ConstantModel, error) {
	var result model.ConstantModel
	err := s.table(ctx, "constants").Where("code = ?", code).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}
	return result, err
}

// UpdateMany implements IConstantStorage.
func (s constantStorage) UpdateMany(ctx context.Context, code constants.GolangServerConstant, value string) error {
	return s.table(ctx, "constants").Where("code = ?", code).Update("value", value).Error
}
