package wstransport

import (
	"github.com/gin-gonic/gin"
	wsbusiness "golang-server/module/core/business/websocket"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
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

func (t WsTransport) CreateNotificationConnection(ctx *gin.Context) {
	userID := ctx.MustGet(constants.XUserID).(string)
	err := t.biz.CreateNotificationConnection(ctx, userID)
	dto.HandleResponse(ctx, nil, err)
}
