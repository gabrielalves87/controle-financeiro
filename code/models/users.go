package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey;autoIncrement; not null; unique" json:"id"`
	Username  string `gorm:"not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"unique; not null" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var Users []User
