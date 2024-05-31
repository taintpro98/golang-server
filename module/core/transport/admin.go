package transport

import (
	"golang-server/module/core/dto"

	"github.com/gin-gonic/gin"
)

func (t *Transport) AdminSyncUsers(ctx *gin.Context) {
	err := t.biz.AdminSyncUsers(ctx)
	dto.HandleResponse(ctx, nil, err)
}

func (t *Transport) AdminCreateMovie(ctx *gin.Context) {
	var data dto.AdminCreateMovieRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	result, err := t.biz.AdminCreateMovie(ctx, data)
	dto.HandleResponse(ctx, result, err)
}

func (t *Transport) AdminCreateSlot(ctx *gin.Context) {
	var data dto.AdminCreateSlotRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	result, err := t.biz.AdminCreateSlot(ctx, data)
	dto.HandleResponse(ctx, result, err)
}

func (t *Transport) AdminCreateRoom(ctx *gin.Context) {
	var data dto.AdminCreateRoomRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	result, err := t.biz.AdminCreateRoom(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
