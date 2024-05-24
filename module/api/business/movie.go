package business

import (
	"context"
	"golang-server/module/api/dto"
	"golang-server/module/api/model"
)

func (b biz) ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error) {
	moviesDB, err := b.movieStorage.List(ctx, dto.FilterMovie{})
	if err != nil {
		return dto.ListMoviesResponse{}, nil, err
	}
	response := dto.ListMoviesResponse{
		Movies: moviesDB,
	}
	count, err := b.movieStorage.Count(ctx, dto.FilterMovie{})
	if err != nil {
		return response, nil, err
	}
	return response, count, nil
}

func (b biz) AdminCreateMovie(ctx context.Context, data dto.AdminCreateMovieRequest) (dto.AdminCreateMovieResponse, error) {
	var response dto.AdminCreateMovieResponse
	dataInsert := model.MovieModel{
		Title:   data.Title,
		Content: data.Content,
	}
	err := b.movieStorage.Insert(ctx, &dataInsert)
	if err != nil {
		return response, err
	}
	response.MovieID = dataInsert.ID
	return response, err
}
