//Routes/Routes.go
package Routes

import (
	"ecapp_cust/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/customer")
	{
		grp1.GET("all_customers", Controllers.GetCustomers)
		grp1.POST("customers", Controllers.CreateCustomer)
		grp1.GET("customers/:id", Controllers.GetCustomerByID)
		grp1.PUT("customers/:id", Controllers.UpdateCustomer)
		grp1.DELETE("customers/:id", Controllers.DeleteCustomer)
	}

	return r
}
