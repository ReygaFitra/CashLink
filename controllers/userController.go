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

func SignUp(c *gin.Context) {
	var body struct{
		Name string
		Username string
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Name: body.Name, Username:body.Username, Email:body.Email, Password: string(hash)}
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Signup",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Signup Success",
	})
}

func Login(c *gin.Context) {
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	err :=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
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

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success",
	})
}

func ViewUser(c *gin.Context) {
	authenticatedUserID, _ := c.Get("authenticatedUserID")
	userID := authenticatedUserID.(uint)

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func FindUserByName(c *gin.Context) {
	username:= c.Param("username")

	var user models.User
	config.DB.First(&user, "username = ?", username)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Id": user.ID,
		"Name": user.Name,
		"Username": user.Username,
		"Email": user.Email,
	})
}

func UpdateUser(c *gin.Context) {
	var body struct{
		Name string
		Username string
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Name: body.Name, Username:body.Username, Email:body.Email, Password: string(hash)}
	result := config.DB.First(&user, "id = ?", c.Param("userID"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found or you don't have permission to update it",
		})
		return
	}

	user.Name = body.Name
	user.Username = body.Username
	user.Email = body.Email
	user.Password = body.Password

	result = config.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User update successfully",
	})
}