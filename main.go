package main

import (
	"go_basic_gorm_gin/config"
	"go_basic_gorm_gin/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("get-token", route.GenerateToken)
		user.POST("register", route.RegisterUser)
	}

	department := router.Group("/department")
	{
		department.GET("/", route.GetDepartment)
		department.GET("/:id", route.GetDepartmentById)
		//department.GET("/department/:name", route.GetDepartmentWithParam)
		department.POST("/", route.PostDepartment)
		department.PUT("/:id", route.UpdateDepartment)
		department.DELETE("/:id", route.DeleteDepartment)
	}
	router.GET("/", getHome)

	position := router.Group("/position")
	{
		position.GET("/", route.GetPosition)
		position.GET("/:id", route.GetPositionById)
		////router.GET("/:name", route.GetDepartmentWithParam)
		position.POST("/", route.PostPosition)
		position.PUT("/:id", route.UpdatePosition)
		position.DELETE("/:id", route.DeletePosition)
	}

	router.Run()
}
func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "ini adalah halaman home",
		"description": "Welcone to Home",
	})
}

//db.Create(Product{
//	Code:  "324",
//	Price: 1200,
//})
//var product Product
//db.First(&product, 1)
//fmt.Println(product)
//
////db.Model(&product).Updates(Product{
////	Code:"A22",
////	Price: 20000,
////
////})
////db.Delete(&product, 1)
