package route

import (
	"go_basic_gorm_gin/config"
	"go_basic_gorm_gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosition(c *gin.Context) {
	position := []model.Positions{}
	config.DB.Find(&position)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data Fineded",
		"data":    position,
	})
}

func GetPositionById(c *gin.Context) {
	id := c.Param("id")
	var position model.Positions
	ress := config.DB.First(&position, "id = ?", id)
	if ress != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Fineded",
			"Data":    position,
		})
	}

}

func PostPosition(c *gin.Context) {
	position := model.Positions{
		Name: c.PostForm("name"),
		Code: c.PostForm("code"),
	}
	config.DB.Create(&position)

	c.JSON(http.StatusCreated, gin.H{
		"data":    position,
		"message": "Input Berhasil",
	})
}

func UpdatePosition(c *gin.Context) {
	id := c.Param("id")
	var position model.Positions

	data := config.DB.First(&position, "id = ?", id)

	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "data not found",
			"message": "data not found",
		})
		return
	}
	config.DB.Model(&position).Where("id = ?", id).Update("Name", c.PostForm("name"))
	config.DB.Model(&position).Where("id = ?", id).Update("Code", c.PostForm("code"))
	c.JSON(200, gin.H{
		"message": "Update Success",
		"data":    position,
	})
}
func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	var position model.Positions
	data := config.DB.First(&position, "id = ?", id)
	if data.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "data not found",
			"message": "data not found",
		})
		return
	}
	config.DB.Delete(&position, id)
	c.JSON(200, gin.H{
		"message": "delete succes",
	})
}
