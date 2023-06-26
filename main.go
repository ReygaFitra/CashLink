package main

import (
	"github.com/ReygaFitra/CashLink.git/config"
	"github.com/ReygaFitra/CashLink.git/delivery"
	"github.com/ReygaFitra/CashLink.git/utils"
)

func init() {
	utils.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	delivery.RunServer()
}