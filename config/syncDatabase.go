package config

import "github.com/ReygaFitra/CashLink.git/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Merchant{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Transfer{})
	DB.AutoMigrate(&models.Payment{})
}