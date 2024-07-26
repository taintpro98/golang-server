package fx_transport

import "github.com/gin-gonic/gin"

type Transport struct {
}

func NewTransport() *Transport {
	return &Transport{}
}

func (t *Transport) Register(ctx *gin.Context) {

}

func (t *Transport) Login(ctx *gin.Context) {

}

func (t *Transport) ListMovies(ctx *gin.Context) {

}
