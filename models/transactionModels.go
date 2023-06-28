package models

import "time"

type Transfer struct {
	Transfer_ID          int64     `gorm:"primaryKey;autoIncrement" json:"tf_id"`
	Transfer_Amount      float64   `gorm:"not null" json:"tf_amount"`
	Transfer_SenderID    int64     `gorm:"not null" json:"tf_sender_id"`
	Transfer_RecipientID int64     `gorm:"not null" json:"tf_recipient_id"`
	Transfer_Description string    `gorm:"not null" json:"tf_description"`
	CreatedAt   time.Time `gorm:"not null" json:"timestamp"`
	Sender   User `gorm:"foreignKey:Transfer_SenderID"`
	Recipient   User `gorm:"foreignKey:Transfer_RecipientID"`
}

type TransferLog struct {
	TransferLog_ID          int64     `gorm:"primaryKey;autoIncrement" json:"tf_log_id"`
	TransferLog_Amount      float64   `gorm:"not null" json:"tf_log_amount"`
	TransferLog_SenderID    int64     `gorm:"not null" json:"tf_log_sender_id"`
	TransferLog_RecipientID int64     `gorm:"not null" json:"tf_log_recipient_id"`
	TransferLog_Description string    `gorm:"not null" json:"tf_log_description"`
	CreatedAt   time.Time `gorm:"not null" json:"timestamp"`
	Transfer                Transfer  `gorm:"foreignKey:TransferLog_ID"`
}

type Payment struct {
	Payment_ID            int64     `gorm:"primaryKey;autoIncrement" json:"pay_id"`
	Payment_UserID        int64     `gorm:"not null" json:"pay_user_id"`
	Payment_MerchantID    int64     `gorm:"not null" json:"pay_merchant_id"`
	Payment_ProductID     int64     `gorm:"not null" json:"pay_product_id"`
	Payment_Amount        float64   `gorm:"not null" json:"pay_amount"`
	CreatedAt     time.Time `gorm:"not null" json:"timestamp"`
	User               User      `gorm:"foreignKey:Payment_UserID"`
	Merchant           Merchant  `gorm:"foreignKey:Payment_MerchantID"`
	Product            Product   `gorm:"foreignKey:Payment_ProductID"`
}

type PaymentLog struct {
	PaymentLog_ID          int64     `gorm:"primaryKey;autoIncrement" json:"pay_log_id"`
	PaymentLog_UserID      int64     `gorm:"not null" json:"pay_log_user_id"`
	PaymentLog_MerchantID  int64     `gorm:"not null" json:"pay_log_merchant_id"`
	PaymentLog_ProductID   int64     `gorm:"not null" json:"pay_log_product_id"`
	PaymentLog_Amount      float64   `gorm:"not null" json:"pay_log_amount"`
	CreatedAt     time.Time `gorm:"not null" json:"timestamp"`
	Payment               Payment   `gorm:"foreignKey:PaymentLog_ID"`
}