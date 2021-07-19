package main

import (
	"ecapp_payment/Controllers"
	"ecapp_payment/Models"
	"ecapp_payment/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//db.DB.

	infrastructure.LoadEnv()
	db := infrastructure.NewDatabase()
	db.DB.AutoMigrate(&Models.Payment{})
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	grp4 := router.Group("/payment")
	{
		grp4.GET("allpayments", Controllers.GetPayment)
		grp4.POST("pay", Controllers.CreatePayment)
		grp4.POST("paying/:id", Controllers.UpdatePayment)
	}
	router.Run(":8003")
}
