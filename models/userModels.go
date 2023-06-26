package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64    `gorm:"primaryKey;autoIncrement" json:"id"` 
	Name     string  `gorm:"not null" json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}