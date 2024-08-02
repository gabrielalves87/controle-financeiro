package models

import "gorm.io/gorm"

type Trasactions struct {
	gorm.Model
	UserId       uint      `gorm:"not null" json:"userId"`
	UserID       User      `gorm:"not null; foreignKey:UserId"`
	CategoriesID uint      `gorm:"not null" json:"categoriesId"`
	Categories   Categorie `gorm:"foreignKey:CategoriesID"`
	Description  string    `gorm:"not null" json:"description"`
	AccountId    uint      `gorm:"not null" json:"accountId"`
	AccountID    Account   `gorm:"not null; foreignKey:AccountId"`
}
