package store

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Label     string
	Embedding []byte `gorm:"type:blob"`
}

type Centroid struct {
	Label string
	Value []byte `gorm:"type:blob"`
}
