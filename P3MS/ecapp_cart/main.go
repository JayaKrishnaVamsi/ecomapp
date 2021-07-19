package main

import (
	"ecapp_cart/Controllers"
	"ecapp_cart/Models"
	"ecapp_cart/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*func allitems(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../ecapp_public/public/home.html")
}*/
func main() {

	router := gin.Default()
	//db.DB.
	infrastructure.LoadEnv()
	db := infrastructure.NewDatabase()
	db.DB.AutoMigrate(&Models.Cart{})
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	})
	grp3 := router.Group("/mycart")
	{
		grp3.GET("allitems", Controllers.GetCart)
		grp3.GET("cart/:id", Controllers.GetCartByID)
		grp3.POST("additem", Controllers.AddCart)

	}
	//http.HandleFunc("/mycart/allitems", allitems)
	router.Run(":8002")
}
