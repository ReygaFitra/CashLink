package models

type User struct {
	ID       int64    `gorm:"primaryKey;autoIncrement" json:"id"` 
	Name     string  `gorm:"not null" json:"name"`
	Username     string  `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}