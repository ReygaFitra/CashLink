package models

type User struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Balance   int64      `gorm:"default:0" json:"balance"`
	Transfers []Transfer `gorm:"foreignKey:SenderID"`
	Payments  []Payment  `gorm:"foreignKey:Payment_UserID"`
}