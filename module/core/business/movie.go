package business

import (
	"context"
	"golang-server/module/core/dto"
)

func (b biz) ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error) {
	moviesDB, err := b.movieStorage.List(ctx, dto.FilterMovie{
		CommonFilter: dto.CommonFilter{
			Select: []string{"id", "title", "content"},
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

// ListMovieSlots implements IBiz.
func (b biz) ListMovieSlots(ctx context.Context, movieID string) (dto.ListMovieSlotsResponse, error) {
	var response dto.ListMovieSlotsResponse
	slots, err := b.slotStorage.List(ctx, dto.FilterSlot{
		MovieID: movieID,
		CommonFilter: dto.CommonFilter{
			Select: []string{"id", "room_id", "start_time"},
		},
	})
	if err != nil {
		return response, err
	}
	response.Slots = slots
	return response, nil
}
