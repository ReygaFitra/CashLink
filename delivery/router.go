package delivery

import (
	"github.com/ReygaFitra/CashLink.git/controllers"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.AuthMiddleware, controllers.Validate)
	
	r.Run()
}