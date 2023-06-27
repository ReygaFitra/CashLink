package delivery

import (
	"github.com/ReygaFitra/CashLink.git/controllers"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	// User Authentication
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout",  controllers.Logout)
	// Find User by Name
	r.GET("/user", controllers.AuthMiddleware, controllers.Validate)
	r.GET("/user/search/:username", controllers.AuthMiddleware, controllers.FindUserByName)
	
	r.Run()
}