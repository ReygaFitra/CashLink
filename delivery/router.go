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
	r.GET("/user", controllers.AuthMiddleware, controllers.ViewUser)
	r.PUT("/user/:id", controllers.AuthMiddleware, controllers.UpdateUser)
	r.GET("/user/search/:username", controllers.AuthMiddleware, controllers.FindUserByName)

	// Merchant Authentication
	r.POST("/signup/merchant", controllers.RegisterMerchant)
	r.POST("/login/merchant", controllers.LoginMerchant)
	r.POST("/merchant/logout",  controllers.Logout)
	r.GET("/merchant", controllers.MerchantMiddleware, controllers.ViewMerchant)
	// Merchant Products
	r.POST("/merchant/product", controllers.MerchantMiddleware, controllers.AddProduct)
	r.PUT("/merchant/product", controllers.MerchantMiddleware, controllers.UpdateProduct)
	r.GET("/merchant/product", controllers.MerchantMiddleware, controllers.ViewProducts)

	// Transaction
	r.POST("user/transaction/transfer/:senderID/:recipientID", controllers.AuthMiddleware, controllers.Transfer)
	r.POST("user/transaction/payment/:userID/:merchantID/:productID", controllers.AuthMiddleware, controllers.Payment)
	r.GET("user/transaction/history/:userID", controllers.AuthMiddleware, controllers.HistoryTransaction)
	
	r.Run()
}