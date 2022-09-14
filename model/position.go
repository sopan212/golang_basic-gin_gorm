package model

import "gorm.io/gorm"

type Positions struct {
	gorm.Model
	Name         string
	Code         string `gorm:"uniq_index"`
	DepartmentID uint
}
