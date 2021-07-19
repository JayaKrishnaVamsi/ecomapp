package main

import (
	"ecapp_prod/Controllers"
	"ecapp_prod/Models"
	"ecapp_prod/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//Config.DB.

	infrastructure.LoadEnv()
	db := infrastructure.NewDatabase()
	db.DB.AutoMigrate(&Models.Product{})
	router.GET("/", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	grp1 := router.Group("/product")
	{
		grp1.GET("allproducts", Controllers.GetProducts)
		grp1.POST("update/:id", Controllers.UpdateProduct)
		grp1.GET("products/:id", Controllers.GetProductByID)

	}
	router.Run(":8001")
}
