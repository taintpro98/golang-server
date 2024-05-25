package business

import (
	"context"
	"golang-server/module/api/dto"
)

func (b biz) ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error) {
	moviesDB, err := b.movieStorage.List(ctx, dto.FilterMovie{
		CommonFilter: dto.CommonFilter{
			Select: []string{"title", "content"},
		},
	})
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
