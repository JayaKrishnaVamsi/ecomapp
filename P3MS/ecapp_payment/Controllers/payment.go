//Controllers/User.go
package Controllers

import (
	"bytes"
	"ecapp_payment/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//var db infrastructure.Database

//GetUsers ... Get all users
func GetPayment(c *gin.Context) {
	var payment []Models.Payment
	err := Models.GetPayment(&payment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

//CreateUser ... Create User
func CreatePayment(c *gin.Context) {
	var payment Models.Payment
	//var cart Models.Cart
	c.BindJSON(&payment)
	err := Models.CreatePayment(&payment)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
	DoPayment(&payment)

	//fmt.Println(payment)
}

func DoPayment(payment *Models.Payment) {
	//fmt.Println(payment)
	po := payment.OId
	pc := payment.CId
	//var result []Models.Cart
	resp, err := http.Get("http://web_cart:8002/mycart/cart/" + strconv.FormatUint(uint64(po), 10))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var responseObject []Models.Cart
	json.Unmarshal(body, &responseObject)

	fmt.Println(responseObject)

	var total uint = 0
	for _, v := range responseObject {
		//db.DB.Raw("SELECT * FROM ecom.product WHERE id = ?", v.PId).Scan(&r1)
		resp, err := http.Get("http://web_prod:8001/product/products/" + strconv.FormatUint(uint64(v.PId), 10))
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var responseObject2 Models.Product
		json.Unmarshal(body, &responseObject2)
		total = total + (uint(responseObject2.Price) * uint(v.PQty))
	}
	fmt.Println("Total Bill")
	fmt.Println(total)
	postBody, _ := json.Marshal(Models.Payment{
		OId:  po,
		CId:  pc,
		Bill: total,
	})
	responseBody := bytes.NewBuffer(postBody)
	url := "http://web_pay:8003/payment/paying/" + strconv.FormatUint(uint64(po), 10)
	//response, err := http.NewRequest("PUT", "http://web_pay:8003/payment/paying/"+strconv.FormatUint(uint64(po), 10), responseBody)
	response, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(response)
	//var responseObject3 []Models.Cart
	//json.Unmarshal(body, &responseObject)
	//db.DB.Exec("UPDATE ecom.payment SET bill = ? WHERE o_id = ?", total, po)
	//fmt.Println("Updated bill in db")

	for _, v := range responseObject {
		//db.DB.Exec("UPDATE ecom.product SET quantity = ? WHERE id = ?", gorm.Expr("quantity - ?", v.PQty), v.PId)
		resp, err := http.Get("http://web_prod:8001/product/products/" + strconv.FormatUint(uint64(v.PId), 10))
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var responseObject3 Models.Product
		json.Unmarshal(body, &responseObject3)
		initquantity := responseObject3.Quantity
		var updatedqty uint
		if initquantity >= v.PQty {
			updatedqty = initquantity - v.PQty
			postBody, _ := json.Marshal(Models.Product{
				Id:       responseObject3.Id,
				Name:     responseObject3.Name,
				Price:    responseObject3.Price,
				Quantity: updatedqty,
			})
			responseBody := bytes.NewBuffer(postBody)
			http.Post("http://web_prod:8001/product/update/"+strconv.FormatUint(uint64(v.PId), 10), "application/json", responseBody)
		}
	}
	fmt.Println("Quantity decremented in db")

}

/*
//GetUserByID ... Get the user by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateUser ... Update the user information*/
func UpdatePayment(c *gin.Context) {
	var payment Models.Payment
	id := c.Params.ByName("id")
	err := Models.GetPaymentByID(&payment, id)
	if err != nil {
		c.JSON(http.StatusNotFound, payment)
	}
	c.BindJSON(&payment)
	err = Models.UpdatePayment(&payment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, payment)
	}
}

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
