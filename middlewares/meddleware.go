package middlewares

import (
	"client-admin/pkg/helperjwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := helperjwt.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, "Invalid Token")
			context.Abort()
			return
		}
		context.Next()
	}
}
