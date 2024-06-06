package transport

import (
	"golang-server/module/core/business"
	"golang-server/module/core/dto"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	biz business.IBiz
	ch  chan string
}

func NewTransport(
	biz business.IBiz,
) *Transport {
	ch := make(chan string)
	return &Transport{
		biz: biz,
		ch:  ch,
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
