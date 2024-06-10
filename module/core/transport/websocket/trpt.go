package wstransport

import (
	wsbusiness "golang-server/module/core/business/websocket"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

type WsTransport struct {
	biz wsbusiness.IWsBusiness
}

func NewWsTransport(
	biz wsbusiness.IWsBusiness,
) WsTransport {
	return WsTransport{
		biz: biz,
	}
}

func (t WsTransport) CreateMsgConnection(ctx *gin.Context) {
	userID := ctx.MustGet(constants.XUserID).(string)
	err := t.biz.CreateMsgConnection(ctx, userID)
	dto.HandleResponse(ctx, nil, err)
}

func (t WsTransport) CreateMsgKafkaConnection(ctx *gin.Context) {
	userID := ctx.MustGet(constants.XUserID).(string)
	err := t.biz.CreateMsgKafkaConnection(ctx, userID)
	dto.HandleResponse(ctx, nil, err)
}
