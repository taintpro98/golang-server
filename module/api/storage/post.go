package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"

	"gorm.io/gorm"
)

type IPostStorage interface {
	FindOne(ctx context.Context, filter dto.FilterPost) (model.PostModel, error)

	List(ctx context.Context, filter dto.FilterPost) ([]model.PostModel, error)

	Insert(ctx context.Context, data *model.PostModel) error

	UpdateMany(ctx context.Context, filter dto.FilterPost, data model.PostModel) error
}

type postStorage struct {
	commonStorage
}

func NewPostStorage(cfg config.DatabaseConfig, db *gorm.DB) IPostStorage {
	return postStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u postStorage) tableName() string {
	return model.PostModel{}.TableName()
}

func (s postStorage) BuildQuery(filter dto.FilterPost) *gorm.DB {
	query := s.table(s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.UserID != "" {
		query = query.Where("user_id = ?", filter.UserID)
	}
	return query
}

// FindOne implements IPostStorage.
func (u postStorage) FindOne(ctx context.Context, filter dto.FilterPost) (model.PostModel, error) {
	var result model.PostModel
	err := u.CFindOne(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

// List implements IPostStorage.
func (u postStorage) List(ctx context.Context, filter dto.FilterPost) ([]model.PostModel, error) {
	var result []model.PostModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(filter),
		Data:         &result,
	})
	return result, err
}

func (u postStorage) Insert(ctx context.Context, data *model.PostModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

// UpdateMany implements IPostStorage.
func (u postStorage) UpdateMany(ctx context.Context, filter dto.FilterPost, data model.PostModel) error {
	return u.CUpdateMany(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
		Query:     u.BuildQuery(filter),
	})
}
