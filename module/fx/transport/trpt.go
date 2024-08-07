package fx_transport

import (
	fx_business "golang-server/module/fx/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz fx_business.IAuthenticateBiz
}

func NewTransport(
	authBiz fx_business.IAuthenticateBiz,
) *Transport {
	return &Transport{
		authBiz: authBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}

func (t *Transport) Register(ctx *gin.Context) {

}

func (t *Transport) Login(ctx *gin.Context) {

}

func (t *Transport) ListMovies(ctx *gin.Context) {

}
