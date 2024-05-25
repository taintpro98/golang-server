package transport

import (
	"golang-server/module/api/dto"

	"github.com/gin-gonic/gin"
)

func (t Transport) ListMovies(ctx *gin.Context) {
	var data dto.ListMoviesRequest
	if err := ctx.ShouldBindQuery(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	movies, total, err := t.biz.ListMovies(ctx, data)
	if err != nil {
		dto.HandleResponse(ctx, nil, err)
	} else {
		limit, offset := data.Paginate.InfoPaginate()
		dto.HandleResponse(ctx, movies, nil, dto.PaginateResponse{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		})
	}
}
