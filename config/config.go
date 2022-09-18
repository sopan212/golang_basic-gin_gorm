package config

import (
	"go_basic_gorm_gin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/golang_basic_gin_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to Connect DB")
	}

	DB.AutoMigrate(&model.Positions{}, &model.Department{}, &model.User{})

	// DB.Create(&model.Positions{
	// 	Name: "Enginer",
	// 	Code: "eng",
	// })
	// DB.Create(&model.Department{
	// 	Name: "HUMAN RESOURCE",
	// 	Code: "hrd",
	// 	Positions: []model.Positions{
	// 		{Name: "General Manager", Code: "GM"},
	// 		{Name: "Manager", Code: "M HRD"},
	// 	},
	// })
	// DB.Create(&model.Department{
	// 	Name: "FINANCE",
	// 	Code: "fin",
	// 	Positions: []model.Positions{
	// 		{Name: "General Manager", Code: "GM"},
	// 		{Name: "Manager", Code: "M FIN"},
	// 	},
	// })
	// DB.Create(&model.User{
	// 	Name:     "Sopan",
	// 	Username: "sopanx",
	// 	Email:    "Sopan@sopa.com",
	// 	Password: "123123",
	// })

}
