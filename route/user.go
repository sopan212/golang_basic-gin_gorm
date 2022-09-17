package route

import (
	"github.com/gin-gonic/gin"
	"go_basic_gorm_gin/model"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  err.Error(),
			"message": "Bad Request",
		})
		c.Abort()
		return
	}

	//hash Password
	if err := user.HashPassword(user.Password);err!= nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":
		})
	}
}
