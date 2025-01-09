package main

import (
	"github.com/gin-gonic/gin"
	"web/controllers"
	"web/db/connection"
)

var addr = "localhost:8080"

func init() {
	connection.ToDB()
}

func main() {

	r := gin.Default()

	r.GET("/", controllers.FindAllProducts)
	r.POST("/createProduct", controllers.CreateProduct)
	r.GET("/findProduct/:id", controllers.FindProductByID)

	if err := r.Run(addr); err != nil {
		panic("failed to connection to address")
	}

}
