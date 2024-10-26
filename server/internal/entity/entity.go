package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint
	Name        string
	Description string
	Price       float64
	ImageURL    string
}
