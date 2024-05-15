package transport

import (
	"golang-server/module/api/dto"
	"golang-server/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetUserPosts(ctx *gin.Context) {
	var data dto.GetUserPostsRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		logger.Error(ctx, err, "fuxkxkxkx")
	}
	result, err := t.biz.GetUserPosts(ctx, data.UserID)
	if err == nil {
		ctx.JSON(http.StatusOK, result)
	}
}

func (t *Transport) GetUserPostByID(ctx *gin.Context) {
	postID := ctx.Param("postID")
	result, err := t.biz.GetUserPostByID(ctx, postID)
	if err == nil {
		ctx.JSON(http.StatusOK, result)
	}
}
