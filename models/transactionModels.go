package models

import "time"

type Transaction struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount      float64   `gorm:"not null" json:"amount"`
	SenderID    int64     `gorm:"not null" json:"sender_id"`
	RecipientID int64     `gorm:"not null" json:"recipient_id"`
	Description string    `gorm:"not null" json:"description"`
	CreatedAt   time.Time `gorm:"not null" json:"timestamp"`
}