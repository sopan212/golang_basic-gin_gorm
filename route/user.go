package route

import (
	"github.com/gin-gonic/gin"
	"go_basic_gorm_gin/config"
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
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	//insert data to table user
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"meesage": "internal server error",
			"error":   insertUser.Error.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"userID":   user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
