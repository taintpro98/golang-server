package transport

import (
	"golang-server/module/api/dto"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetMovieSlotInfo(ctx *gin.Context) {
	slotID := ctx.Param("slotID")
	result, err := t.biz.GetMovieSlotInfo(ctx, slotID)
	dto.HandleResponse(ctx, result, err)
}

func (t *Transport) ReserveSeats(ctx *gin.Context) {
	var data dto.ReserveSeatsRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	result, err := t.biz.ReserveSeats(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
