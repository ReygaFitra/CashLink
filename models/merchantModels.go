package models

type Merchant struct {
	Merchant_ID       int64     `gorm:"primaryKey;autoIncrement" json:"merchant_id"`
	Merchant_Name     string    `gorm:"not null" json:"merchant_name"`
	Merchant_Email    string    `gorm:"unique" json:"merchant_email"`
	Merchant_Password string    `gorm:"not null" json:"merchant_password"`
	Merchant_Address  string    `gorm:"not null" json:"merchant_address"`
	Products          []Product `gorm:"foreignKey:MerchantID"`
	Payments          []Payment `gorm:"foreignKey:Payment_MerchantID"`
}

type Product struct {
	Product_ID    int64     `gorm:"primaryKey;autoIncrement" json:"product_id"`
	Product_Name  string    `gorm:"not null" json:"product_name"`
	Product_Price float64   `gorm:"not null" json:"product_price"`
	MerchantID    int64     `gorm:"not null" json:"merchant_id"`
	Merchant      Merchant  `gorm:"foreignKey:MerchantID"`
	Payments      []Payment `gorm:"foreignKey:Payment_ProductID"`
}
