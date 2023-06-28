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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Saldo pengirim tidak cukup"})
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
