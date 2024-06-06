package transport

import (
	"github.com/gin-gonic/gin"
	"golang-server/module/core/dto"
	"io"
	"net/http"
)

func (t *Transport) HandleEventStreamPost(ctx *gin.Context) {
	var request dto.EventStreamRequest
	if err := ctx.ShouldBind(&request); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}

	t.ch <- request.Message
	ctx.JSON(http.StatusOK, "done")
}

func (t *Transport) HandleEventStreamGet(ctx *gin.Context) {
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-t.ch; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}

func (t *Transport) CreateEventStreamConnection(ctx *gin.Context) {
	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-t.ch; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}
