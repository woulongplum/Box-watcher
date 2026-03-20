package model

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	Name string `gorm:"not null"`
	URL string	`gorm:"uniqueIndex;not null"`
	Price int
	InStock bool	`gorm:"default:false"`
	Source string
	LastNotifiedAt *time.Time
}
