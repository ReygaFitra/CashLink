package config

import "github.com/ReygaFitra/CashLink.git/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Merchant{})
}