package route

import (
	"go_basic_gorm_gin/auth"
	"go_basic_gorm_gin/config"
	"go_basic_gorm_gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	//check email
	checkEmail := config.DB.Where("email = ?", request.Email).First(&user)
	if checkEmail.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   checkEmail.Error.Error(),
		})
		c.Abort()
		return
	}

	//check password
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "An Authorized",
			"error":   credentialError.Error(),
		})
		c.Abort()
		return
	}

	//Genertate Token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	//Respon Token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
