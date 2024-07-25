package storage

import (
	"context"
	"golang-server/config"
	"golang-server/module/core/dto"
	"golang-server/module/core/model"

	"gorm.io/gorm"
)

type IMovieStorage interface {
	Count(ctx context.Context, filter dto.FilterMovie) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterMovie) (model.MovieModel, error)

	List(ctx context.Context, filter dto.FilterMovie) ([]model.MovieModel, error)

	Insert(ctx context.Context, data *model.MovieModel) error
}

type movieStorage struct {
	commonStorage
}

func NewMovieStorage(cfg config.DatabaseConfig, db *gorm.DB) IMovieStorage {
	return movieStorage{
		commonStorage: commonStorage{
			db:       db,
			configDb: cfg,
		},
	}
}

func (u movieStorage) tableName() string {
	return model.MovieModel{}.TableName()
}

func (s movieStorage) BuildQuery(ctx context.Context, filter dto.FilterMovie) *gorm.DB {
	query := s.table(ctx, s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	return query
}

func (u movieStorage) FindOne(ctx context.Context, filter dto.FilterMovie) (model.MovieModel, error) {
	var result model.MovieModel
	err := u.CFindOne(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

// Count implements IMovieStorage.
func (u movieStorage) Count(ctx context.Context, filter dto.FilterMovie) (*int64, error) {
	return u.CCount(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
	})
}

// List implements IPostStorage.
func (u movieStorage) List(ctx context.Context, filter dto.FilterMovie) ([]model.MovieModel, error) {
	var result []model.MovieModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonStorageParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u movieStorage) Insert(ctx context.Context, data *model.MovieModel) error {
	return u.CInsert(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

// UpdateMany implements IPostStorage.
func (u movieStorage) UpdateMany(ctx context.Context, filter dto.FilterMovie, data model.MovieModel) error {
	return u.CUpdateMany(ctx, CommonStorageParams{
		TableName: u.tableName(),
		Data:      data,
		Query:     u.BuildQuery(ctx, filter),
	})
}
