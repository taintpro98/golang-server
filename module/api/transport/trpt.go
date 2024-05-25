package transport

import (
	"golang-server/module/api/business"
	"golang-server/module/api/dto"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	biz business.IBiz
}

func NewTransport(
	biz business.IBiz,
) *Transport {
	return &Transport{
		biz: biz,
	}
}

func (t *Transport) Register(ctx *gin.Context) {
	var data dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	result, err := t.biz.Register(ctx, data)
	dto.HandleResponse(ctx, result, err)
}
