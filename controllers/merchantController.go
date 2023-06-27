package controllers

import (
	"net/http"

	"github.com/ReygaFitra/CashLink.git/config"
	"github.com/ReygaFitra/CashLink.git/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterMerchant(c *gin.Context) {
	var body struct {
		Merchant_Name     string
		Merchant_Email    string
		Merchant_Password string
		Merchant_Address  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Merchant_Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	merchant := models.Merchant{Merchant_Name: body.Merchant_Name, Merchant_Email:body.Merchant_Email,  Merchant_Password: string(hash), Merchant_Address: body.Merchant_Address}
	result := config.DB.Create(&merchant)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Register Merchant",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Merchant Registered",
	})
}