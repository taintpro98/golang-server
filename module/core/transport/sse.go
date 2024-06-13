package transport

import (
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SSEWriter is a custom io.Writer for sending SSE data
type SSEWriter struct {
	writer http.ResponseWriter
}

// NewSSEWriter creates a new SSEWriter
func NewSSEWriter(writer http.ResponseWriter) *SSEWriter {
	return &SSEWriter{
		writer: writer,
	}
}

// WriteEvent writes SSE event data
func (w *SSEWriter) WriteEvent(event, data string) {
	w.writer.Write([]byte("event: " + event + "\n"))
	w.writer.Write([]byte("data: " + data + "\n\n"))
}

// Flush pushes SSE data to the client
func (w *SSEWriter) Flush() {
	w.writer.(http.Flusher).Flush()
}

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
	userID := ctx.MustGet(constants.XUserID).(string)

	pubsub := t.biz.HandleEventStreamConnection(ctx, userID)
	// if err != nil {
	// 	dto.HandleResponse(ctx, nil, err)
	// 	return
	// }
	defer pubsub.Close()
	// Set headers for SSE
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	writer := NewSSEWriter(ctx.Writer)

	ctx.Stream(func(w io.Writer) bool {
		// Receive messages from Redis and send them to the client via SSE
		for {
			msg, err := pubsub.ReceiveMessage(ctx.Request.Context())
			if err != nil {
				writer.WriteEvent("error", err.Error())
				writer.Flush()
				break
			}
			writer.WriteEvent("message", msg.Payload)
			writer.Flush()
		}
		return false
	})
}
