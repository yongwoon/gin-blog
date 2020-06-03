package entities

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:"size:255"`
	Contents string `gorm:"type:text"`
}
