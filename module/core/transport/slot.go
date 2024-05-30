package transport

import (
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetMovieSlotInfo(ctx *gin.Context) {
	slotID := ctx.Param("slotID")
	result, err := t.biz.GetMovieSlotInfo(ctx, slotID)
	dto.HandleResponse(ctx, result, err)
}

func (t *Transport) ReserveSeats(ctx *gin.Context) {
	slotID := ctx.Param("slotID")
	var data dto.ReserveSeatsRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	data.UserID = ctx.MustGet(constants.XUserID).(string)
	result, err := t.biz.ReserveSeats(ctx, slotID, data)
	dto.HandleResponse(ctx, result, err)
}
