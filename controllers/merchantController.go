package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/ReygaFitra/CashLink.git/config"
	"github.com/ReygaFitra/CashLink.git/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func LoginMerchant(c *gin.Context) {
	var body struct {
		Merchant_Email    string
		Merchant_Password string
	} 

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var merchant models.Merchant
	config.DB.First(&merchant, "merchant_email = ?", body.Merchant_Email)
	if merchant.Merchant_ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	err :=bcrypt.CompareHashAndPassword([]byte(merchant.Merchant_Password), []byte(body.Merchant_Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": merchant.Merchant_ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
	})
}

func MerchantValidate(c *gin.Context) {
	merchant, _ := c.Get("merchant")

	c.JSON(http.StatusOK, gin.H{
		"merchant": merchant,
	})
}

func AddProduct(c *gin.Context) {
	var body struct {
		Product_Name  string
		Product_Price float64
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	merchant, exists := c.Get("merchant")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get merchant information",
		})
		return
	}

	merchantID, ok := merchant.(models.Merchant)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	product := models.Product{
		Product_Name:  body.Product_Name,
		Product_Price: body.Product_Price,
		MerchantID:    merchantID.Merchant_ID,
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create Product",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Add New Product Successfully",
	})
}

func UpdateProduct(c *gin.Context) {
	var body struct {
		Product_ID     int64
		Product_Name   string
		Product_Price  float64
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	merchant, exists := c.Get("merchant")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get merchant information",
		})
		return
	}

	merchantID, ok := merchant.(models.Merchant)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid merchant ID",
		})
		return
	}

	var product models.Product
	result := config.DB.First(&product, "product_id = ? AND merchant_id = ?", body.Product_ID, merchantID.Merchant_ID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found or you don't have permission to update it",
		})
		return
	}

	product.Product_Name = body.Product_Name
	product.Product_Price = body.Product_Price

	result = config.DB.Save(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})
}
