package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model
	Code     string  `gorm:"column:code;unique" json:"code"`
	Name     string  `gorm:"column:name;unique" json:"name"`
	Symbol   string  `gorm:"column:symbol" json:"symbol"`
	Exchange float64 `gorm:"column:exchange" json:"exchange"`
}
