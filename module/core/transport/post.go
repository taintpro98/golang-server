package transport

import (
	"github.com/gin-gonic/gin"
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
)

func (t *Transport) CreatePost(ctx *gin.Context) {
	var data dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, data, err)
		return
	}
	userID := ctx.MustGet(constants.XUserID).(string)
	result, err := t.biz.CreatePost(ctx, userID, data)
	dto.HandleResponse(ctx, result, err)
}
