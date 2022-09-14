package route

import (
	"github.com/gin-gonic/gin"
	config "go_basic_gorm_gin/config"
	"go_basic_gorm_gin/model"
	"net/http"
)

func GetDepartment(c *gin.Context) {
	departments := []model.Department{}
	config.DB.Preload("Positions").Find(&departments)
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Departments",
		"data":    departments,
	})
}

//func GetDepartmentWithParam(c *gin.Context) {
//	name := c.Param("name")
//	c.JSON(http.StatusOK, gin.H{
//		"message":     "hallo " + name + " ini adalah halaman home",
//		"description": "Welcone to Home",
//	})
//}
func GetDepartmentById(c *gin.Context) {
	id := c.Param("id")
	var department model.Department
	ress := config.DB.Preload("Positions").First(&department, "id = ?", id)
	if ress.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message":     "Data Not Found",
			"Description": "Data Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    ress,
		"message": "success",
	})
}

func PostDepartment(c *gin.Context) {
	var department = model.Department{
		Name: c.PostForm("name"),
		Code: c.PostForm("code"),
	}
	config.DB.Create(&department)
	//if result.Error != nil {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "Insert Not Founds",
	//	})
	//	return
	//}
	c.JSON(http.StatusCreated, gin.H{
		"data":    department,
		"message": "Inserted Success",
	})
}

func UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var department model.Department
	r := config.DB.First(&department, "id = ?", id)
	if r.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
			"Data":    "Data Not Found",
		})
		return
	}
	config.DB.Model(&department).Where("id = ?", id).Update("Name", c.PostForm("name"))
	config.DB.Model(&department).Where("id = ?", id).Update("Code", c.PostForm("code"))

	c.JSON(200, gin.H{
		"message": "Update Succes",
		"Data":    department,
	})
}
func DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var department model.Department
	re := config.DB.First(&department, "id = ?", id)
	if re != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page Not Found",
			"data":    "Page Not Found",
		})
		return
	}
	config.DB.Delete(&department, id)

	c.JSON(200, gin.H{
		"message": "Dellete Success",
	})
}
