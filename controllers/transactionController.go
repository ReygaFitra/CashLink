package controllers

import (
	"net/http"
	"strconv"

	"github.com/ReygaFitra/CashLink.git/config"
	"github.com/ReygaFitra/CashLink.git/models"
	"github.com/gin-gonic/gin"
)

func Transfer(c *gin.Context) {
	senderIDStr := c.Param("senderID")
	senderID, err := strconv.ParseUint(senderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid sender ID",
		})
		return
	}

	recipientIDStr := c.Param("recipientID")
	recipientID, err := strconv.ParseUint(recipientIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid recipient ID",
		})
		return
	}

	var body struct {
		Amount float64
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	var sender, recipient models.User
	if err := config.DB.First(&sender, senderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find sender",
		})
		return
	}
	if err := config.DB.First(&recipient, recipientID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find recipient",
		})
		return
	}

	transfer := models.Transfer{
		SenderID:   uint(senderID),
		ReceiverID: uint(recipientID),
		Amount:     body.Amount,
	}

	if sender.Balance < int64(transfer.Amount) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	result := config.DB.Create(&transfer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process transfer",
		})
		return
	}

	amount := int64(body.Amount)
	sender.Balance -= amount
	if err := config.DB.Save(&sender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update sender's balance",
		})
		return
	}

	recipient.Balance += amount
	if err := config.DB.Save(&recipient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update recipient's balance",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transfer successful",
	})
}

func Payment(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	merchantIDStr := c.Param("merchantID")
	merchantID, err := strconv.ParseUint(merchantIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	productIDStr := c.Param("productID")
	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	var body struct {
		Payment_Amount float64 `json:"payment_amount"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	var user models.User
	var merchant models.Merchant
	var product models.Product
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find user",
		})
		return
	}

	if err := config.DB.First(&merchant, merchantID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find merchant",
		})
		return
	}

	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to find product",
		})
		return
	}

	if body.Payment_Amount != product.Product_Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment amount"})
		return
	}

	payment := models.Payment{
		Payment_UserID:     uint(userID),
		Payment_MerchantID: uint(merchantID),
		Payment_ProductID:  uint(productID),
		Payment_Amount:     body.Payment_Amount,
	}

	if user.Balance < int64(payment.Payment_Amount) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	result := config.DB.Create(&payment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process payment",
		})
		return
	}

	amount := int64(body.Payment_Amount)
	user.Balance -= amount
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user's balance",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment successful",
	})
}
