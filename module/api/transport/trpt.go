package transport

import (
	"github.com/gin-gonic/gin"
	"golang-server/module/api/business"
	"golang-server/module/api/dto"
	"net/http"
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

func (t *Transport) CreateUser(ctx *gin.Context) {
	var data dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {

	}
	result, err := t.biz.CreateUser(ctx, data)
	if err != nil {

	} else {
		ctx.JSON(http.StatusCreated, result)
	}
}

func (t *Transport) GetSports(ctx *gin.Context) {
	_ = t.biz.GetSports(ctx)
}
