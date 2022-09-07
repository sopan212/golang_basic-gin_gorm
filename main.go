package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
type Department struct {
	gorm.Model
	Name string
	Code string `gorm:"uniq_index"`
}
type Position struct {
	gorm.Model
	Name string
	Code string `gorm:"uniq_index"`
}

var DB *gorm.DB

func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/golang_basic_gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to Connect DB")
	}
	DB.AutoMigrate(&Department{}, &Position{})
	DB.Create(&Department{
		Name: "Human Resource",
		Code: "HR",
	})
	DB.Create(&Department{
		Name: "Finance",
		Code: "Fin",
	})

	router := gin.Default()

	router.GET("/", getHome)
	router.GET("/department", getDepartment)
	//router.GET("/department/:id", getDepartmentById)
	//router.GET("/department/:name", getDepartmentWithParam)
	router.POST("/department", PostDepartment)

	router.Run()
}
func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":     "ini adalah halaman home",
		"description": "Welcone to Home",
	})
}
func getDepartment(c *gin.Context) {
	departments := []Department{}
	DB.Find(&departments)
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Departments",
		"data":    departments,
	})
}
func getDepartmentWithParam(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message":     "hallo " + name + " ini adalah halaman home",
		"description": "Welcone to Home",
	})
}
func getDepartmentById(c *gin.Context) {
	id := c.Param("id")
	var department Department
	ress := DB.First(&department, id)
	if ress.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message":     "Data Not Found",
			"Description": "Data Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": ress,
	})
	log.Println(ress)
}
func PostDepartment(c *gin.Context) {
	name := c.PostForm("name")
	code := c.PostForm("code")
	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"code":    code,
		"message": "Inserted Success",
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
