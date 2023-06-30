package models

type Transfer struct {
	ID         uint    `gorm:"primaryKey"`
	SenderID   uint    `gorm:"not null" json:"sender_id"`
	ReceiverID uint    `gorm:"not null" json:"receiver_id"`
	Amount     float64 `gorm:"not null" json:"amount"`
	Sender     User    `gorm:"foreignKey:SenderID"`
	Receiver   User    `gorm:"foreignKey:ReceiverID"`
}

type Payment struct {
	Payment_ID         uint     `gorm:"primaryKey;autoIncrement"`
	Payment_UserID     uint     `gorm:"not null" json:"payment_userid"`
	Payment_MerchantID uint     `gorm:"not null" json:"payment_merchantid"`
	Payment_ProductID  uint     `gorm:"not null" json:"payment_productid"`
	Payment_Amount     float64  `gorm:"not null"`
	User               User     `gorm:"foreignKey:Payment_UserID"`
	Merchant           Merchant `gorm:"foreignKey:Payment_MerchantID"`
	Product            Product  `gorm:"foreignKey:Payment_ProductID"`
}
