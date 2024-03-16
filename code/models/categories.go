package models

import (
	"gorm.io/gorm"
)

type Categorie struct {
	gorm.Model
	// ID          uint   `gorm:"primaryKey;autoIncrement; not null; unique" json:"id"`
	UserId      uint   `gorm:"not null" json:"user_id"`
	UserID      User   `gorm:"not null; foreignKey:UserId"`
	Title       string `gorm:"not null" json:"title"`
	Type        string `gorm:"not null" json:"type"`
	Description string `gorm:"not null" json:"description"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
}
