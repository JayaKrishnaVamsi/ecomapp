//Controllers/User.go
package Controllers

import (
	"fmt"
	"net/http"

	"ecapp_cart/Models"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetCart(c *gin.Context) {
	var cart []Models.Cart
	err := Models.GetCart(&cart)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func AddCart(c *gin.Context) {
	var cart Models.Cart
	c.BindJSON(&cart)
	err := Models.AddCart(&cart)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
	//var pay Models.Payment
	//pay.CId = cart.CId
	//pay.OId = cart.OrderId
	//var result []Models.Cart
	//ordid := cart.OrderId

	//pay.Bill = cart.PQty * result.Price
	//fmt.Println("Total Bill:")
	//fmt.Print(ordid)
	//fmt.Println(result)
	//err = Models.CreatePayment(&pay, &cart)
}

//GetUserByID ... Get the user by id
func GetCartByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var cart []Models.Cart
	err := Models.GetCartByID(&cart, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

/*
//UpdateUser ... Update the user information
func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
*/
//DeleteUser ... Delete the user
/*func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}*/
