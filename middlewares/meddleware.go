package middlewares

import (
	"client-admin/pkg/helperjwt"
	"fmt"

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

		fmt.Print("testing auth   \t")
		err := helperjwt.ValidateToken(tokenString)
		fmt.Println(err)
		if err != nil {
			context.JSON(401, gin.H{
				"err":    err.Error(),
				"reason": "failed to handle",
			})
			context.Abort()
			return
		}
		context.Next()
	}

}

func AuthForUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err := helperjwt.ValidateTokenUserRole(tokenString)
		fmt.Println(err)
		if err != nil {
			context.JSON(401, gin.H{
				"err":    err.Error(),
				"reason": "failed to handle",
			})
			context.Abort()
			return
		}
		context.Next()
	}
}
