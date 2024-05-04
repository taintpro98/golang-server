package transport

import (
	"github.com/gin-gonic/gin"
	"golang-server/module/api/business"
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

func (t *Transport) GetSports(ctx *gin.Context) {
	_ = t.biz.GetSports(ctx)
}
