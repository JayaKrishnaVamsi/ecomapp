package main

import (
	"ecapp_cust/Controllers"
	"ecapp_cust/Models"
	"ecapp_cust/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	//Config.DB.

	infrastructure.LoadEnv()
	db := infrastructure.NewDatabase()
	db.DB.AutoMigrate(&Models.Customer{})
	router.GET("/", func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	grp1 := router.Group("/customer")
	{
		grp1.GET("all_customers", Controllers.GetCustomers)
		grp1.POST("customers", Controllers.CreateCustomer)
		grp1.GET("customers/:id", Controllers.GetCustomerByID)
		grp1.PUT("customers/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customers/:id", Controllers.DeleteCustomer)
	}
	router.Run(":8000")
}
