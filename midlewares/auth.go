package midlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/nacl/auth"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Reauest need access token",
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Erorr(),
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
