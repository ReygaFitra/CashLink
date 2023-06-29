package delivery

import (
	"github.com/ReygaFitra/CashLink.git/controllers"
	"github.com/ReygaFitra/CashLink.git/middleware"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	// User Authentication
	r.POST("/signup/user", controllers.SignUp)
	r.POST("/login/user", controllers.Login)
	r.POST("/user/logout",  controllers.Logout)
	r.GET("/user/:userID", middleware.AuthMiddleware, controllers.ViewUser)
	r.PUT("/user/:userID", middleware.AuthMiddleware, controllers.UpdateUser)
	r.GET("/user/search/:userID/:username", middleware.AuthMiddleware, controllers.FindUserByName)

	// Merchant Authentication
	r.POST("/signup/merchant", controllers.RegisterMerchant)
	r.POST("/login/merchant", controllers.LoginMerchant)
	r.POST("/merchant/logout",  controllers.Logout)
	r.GET("/merchant/:merchantID", middleware.MerchantMiddleware, controllers.ViewMerchant)
	// Merchant Products
	r.POST("/merchant/product/:merchantID", middleware.MerchantMiddleware, controllers.AddProduct)
	r.PUT("/merchant/product/:merchantID", middleware.MerchantMiddleware, controllers.UpdateProduct)
	r.GET("/merchant/product/:merchantID", middleware.MerchantMiddleware, controllers.ViewProducts)

	// Transaction
	r.POST("user/transaction/transfer/:userID/:recipientID", middleware.AuthMiddleware, controllers.Transfer)
	r.POST("user/transaction/payment/:userID/:merchantID/:productID", middleware.AuthMiddleware, controllers.Payment)
	r.GET("user/transaction/history/:userID", middleware.AuthMiddleware, controllers.HistoryTransaction)
	
	r.Run()
}