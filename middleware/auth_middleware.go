package middleware

import (
	"golang-server/module/core/dto"
	"golang-server/pkg/constants"
	"golang-server/pkg/e"
	"golang-server/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtMaker token.IJWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearLength := len("Bearer ")
		if len(authHeader) < bearLength {
			dto.AbortJSON(c, e.ErrUnauthorized)
			return
		}
		tokenString := authHeader[bearLength:]
		payload, err := jwtMaker.VerifyToken(c, tokenString)
		if err != nil {
			dto.AbortJSON(c, e.ErrUnauthorized)
			return
		}
		c.Set(constants.XUserID, payload.(*dto.UserPayload).UserID)
	}
}
