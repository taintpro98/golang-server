package repository

import (
	"context"
	"fmt"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"
	"golang-server/pkg/constants"
	"golang-server/pkg/logger"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Count(ctx context.Context, filter dto.FilterUser) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterUser) (model.UserModel, error)

	List(ctx context.Context, filter dto.FilterUser) ([]model.UserModel, error)

	Insert(ctx context.Context, data *model.UserModel) error

	InsertBatch(ctx context.Context, data *[]model.UserModel) error

	TxInsertMUsers(ctx context.Context, num int, data *[]model.UserModel) error
}

type userRepository struct {
	commonRepository
}

func NewUserRepository(cfg config.DatabaseConfig, db *gorm.DB) IUserRepository {
	return userRepository{
		commonRepository: commonRepository{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u userRepository) tableName() string {
	return model.UserModel{}.TableName()
}

func (s userRepository) BuildQuery(ctx context.Context, filter dto.FilterUser) *gorm.DB {
	query := s.table(ctx, s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Phone != "" {
		query = query.Where("phone = ?", filter.Phone)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	return query
}

// Count implements IUserStorage.
func (u userRepository) Count(ctx context.Context, filter dto.FilterUser) (*int64, error) {
	return u.CCount(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
	})
}

func (u userRepository) FindOne(ctx context.Context, filter dto.FilterUser) (model.UserModel, error) {
	var result model.UserModel
	err := u.CFindOne(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u userRepository) Insert(ctx context.Context, data *model.UserModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

// List implements IUserStorage.
func (u userRepository) List(ctx context.Context, filter dto.FilterUser) ([]model.UserModel, error) {
	var result []model.UserModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

// InsertBatch implements IUserStorage.
func (u userRepository) InsertBatch(ctx context.Context, data *[]model.UserModel) error {
	return u.CInsertBatch(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

func (u userRepository) TxInsertMUsers(ctx context.Context, num int, data *[]model.UserModel) error {
	// Create a transaction
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			logger.Info(ctx, "rollback TxInsertMUsers", logger.LogField{
				Key:   "data",
				Value: data,
			})
			tx.Rollback()
		}
	}()
	err := tx.Create(data).Error
	if err != nil {
		logger.Error(ctx, err, "TxInsertMUsers create users error")
		return err
	}
	err = tx.Model(&model.ConstantModel{}).Where("code = 'users_num'").Update("value", fmt.Sprintf("%d", num+constants.MBatchSize)).Error
	if err != nil {
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		logger.Error(ctx, err, "TxInsertMUsers Error when commit transaction")
		return err
	}
	return nil
}
