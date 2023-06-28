package models

type User struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Amount    int64      `json:"amount"`
	Transfers []Transfer `gorm:"foreignKey:Transfer_SenderID"`
	Receivers []Transfer `gorm:"foreignKey:Transfer_RecipientID"`
	Payments  []Payment  `gorm:"foreignKey:Payment_UserID"`
}