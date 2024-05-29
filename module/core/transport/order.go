package transport

import (
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

func (t *Transport) CreateOrder(ctx *gin.Context) {
	var data dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	userID := ctx.GetHeader(constants.XUserID)
	result, err := t.biz.CreateOrder(ctx, userID, data)
	dto.HandleResponse(ctx, result, err)
}
