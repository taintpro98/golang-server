package middleware

import (
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
	"golang-server/token"

	"github.com/gin-gonic/gin"
)

type WebsocketSecureParam struct {
	Token string `form:"token,omitempty"`
}

func readToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	bearLength := len("Bearer ")
	if len(authHeader) < bearLength {
		var ws WebsocketSecureParam
		if err := c.ShouldBindQuery(&ws); err != nil {
			return "", e.ErrUnauthorized
		}
		return ws.Token, nil
	}
	return authHeader[bearLength:], nil
}

func AuthMiddleware(jwtMaker token.IJWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := readToken(c)
		if err != nil {
			dto.AbortJSON(c, err)
			return
		}
		payload, err := jwtMaker.VerifyToken(c, tokenString)
		if err != nil {
			dto.AbortJSON(c, e.ErrUnauthorized)
			return
		}
		c.Set(constants.XUserID, payload.(*dto.UserPayload).UserID)
	}
}
