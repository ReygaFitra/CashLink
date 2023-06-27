package delivery

import (
	"github.com/ReygaFitra/CashLink.git/controllers"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	// User Authentication
	r.POST("/signup/user", controllers.SignUp)
	r.POST("/login/user", controllers.Login)
	r.POST("/user/logout",  controllers.Logout)
	// Find User by Name
	r.GET("/user", controllers.AuthMiddleware, controllers.Validate)
	r.GET("/user/search/:username", controllers.AuthMiddleware, controllers.FindUserByName)

	// Merchant Authentication
	r.POST("/signup/merchant", controllers.RegisterMerchant)
	r.POST("/login/merchant", controllers.LoginMerchant)
	r.POST("/merchant/logout",  controllers.Logout)

	r.GET("/merchant", controllers.MerchantMiddleware, controllers.MerchantValidate)
	r.POST("/merchant/product", controllers.MerchantMiddleware, controllers.AddProduct)
	r.PUT("/merchant/product", controllers.MerchantMiddleware, controllers.UpdateProduct)
	
	r.Run()
}