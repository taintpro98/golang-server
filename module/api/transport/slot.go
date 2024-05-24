package transport

import (
	"golang-server/module/api/dto"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetMovieSlotInfo(ctx *gin.Context) {
	var data dto.GetMovieSlotInfoRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
	}
	result, err := t.biz.GetMovieSlotInfo(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
