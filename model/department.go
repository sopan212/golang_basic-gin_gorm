package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name      string
	Code      string `gorm:"uniq_index"`
	Positions []Positions
}
