package route

import (
	"github.com/gin-gonic/gin"
	"go_basic_gorm_gin/config"
	"go_basic_gorm_gin/model"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	//validation
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		c.Abort()
		return
	}

	//hass password
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"err":     err.Error(),
		})
	}

	//isert data to table user
	insertUser := config.DB.Create(&user)
	if insertUser.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"err":     insertUser.Error.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"UserID":   user.ID,
		"Email":    user.Email,
		"Username": user.Username,
	})
}
