package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey;autoIncrement; not null; unique" json:"id"`
	UserId       uint      `gorm:"not null" json:"userId"`
	UserID       User      `gorm:"not null; foreignKey:UserId"`
	CategoriesID uint      `gorm:"not null" json:"categoriesId"`
	Categories   Categorie `gorm:"foreignKey:CategoriesID"`
	Title        string    `gorm:"not null" json:"title"`
	Type         string    `gorm:"not null" json:"type"`
	Description  string    `gorm:"not null" json:"description"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var Accounts []Account
